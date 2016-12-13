package transaction

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestNewManager(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)

	connection, err := gorm.Open("mysql", db)
	assert.NoError(t, err)

	manager := NewManager(connection)
	if assert.NotNil(t, manager) {
		assert.Equal(t, uint64(0), manager.counter)
	}
}

func TestStart(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectRollback()

	connection, err := gorm.Open("mysql", db)
	assert.NoError(t, err)

	manager := NewManager(connection)

	old := DEFAULT_TRANSACTION_TIMEOUT
	DEFAULT_TRANSACTION_TIMEOUT = 15 * time.Millisecond

	manager.Start("test")
	time.Sleep(100 * time.Millisecond)

	_, err = manager.Checkout("test-1")
	assert.Error(t, err)

	DEFAULT_TRANSACTION_TIMEOUT = old
}

func TestCheckout(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectCommit()

	connection, err := gorm.Open("mysql", db)
	assert.NoError(t, err)

	manager := NewManager(connection)
	transaction, _ := manager.Start("test")

	_, err = manager.Checkout("test-1")
	assert.NoError(t, err)

	assert.NoError(t, transaction.Commit())

	_, err = manager.Checkout("test-1")
	assert.Error(t, err)
}
