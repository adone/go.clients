package transaction

import (
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestNew(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectRollback()

	connection, err := gorm.Open("mysql", db)
	assert.NoError(t, err)

	old := DEFAULT_TRANSACTION_TIMEOUT
	DEFAULT_TRANSACTION_TIMEOUT = 15 * time.Millisecond

	transaction := New(connection)
	transaction.Start()
	time.Sleep(20 * time.Millisecond)
	assert.NoError(t, transaction.Commit())

	DEFAULT_TRANSACTION_TIMEOUT = old
}

func TestCommit(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectCommit()

	connection, err := gorm.Open("mysql", db)
	assert.NoError(t, err)

	transaction := New(connection)
	transaction.Start()
	assert.NoError(t, transaction.Commit())
}

func TestRollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectRollback()

	connection, err := gorm.Open("mysql", db)
	assert.NoError(t, err)

	transaction := New(connection)
	transaction.Start()
	assert.NoError(t, transaction.Rollback())
}
