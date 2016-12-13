package transaction

import (
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

var DEFAULT_TRANSACTION_TIMEOUT = time.Minute

func New(connection *gorm.DB) *Adapter {
	transaction := new(Adapter)
	transaction.origin = connection.Begin()
	transaction.commit = make(chan struct{})
	transaction.rollback = make(chan struct{})
	transaction.result = make(chan error, 1)

	return transaction
}

type Adapter struct {
	key      string
	origin   *gorm.DB
	guard    sync.Once
	commit   chan struct{}
	rollback chan struct{}
	result   chan error
}

func (transaction *Adapter) Start() {
	go transaction.wait(transaction.result)
}

func (transaction *Adapter) Rollback() (err error) {
	select {
	case err = <-transaction.result:
	default:
		transaction.guard.Do(func() {
			transaction.rollback <- struct{}{}
			err = <-transaction.result
		})
	}

	return
}

func (transaction *Adapter) Commit() (err error) {
	select {
	case err = <-transaction.result:
	default:
		transaction.guard.Do(func() {
			transaction.commit <- struct{}{}
			err = <-transaction.result
		})
	}

	return
}

func (transaction Adapter) Key() string {
	return transaction.key
}

func (transaction *Adapter) wait(result chan error) {
	defer transaction.guard.Do(func() {
		close(result)
		close(transaction.commit)
		close(transaction.rollback)
	})

	select {
	case <-time.After(DEFAULT_TRANSACTION_TIMEOUT):
		result <- transaction.origin.Rollback().Error
	case <-transaction.commit:
		result <- transaction.origin.Commit().Error
	case <-transaction.rollback:
		result <- transaction.origin.Rollback().Error
	}
}
