package transaction

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/pg.v5"
)

func TestNew(t *testing.T) {
	db := pg.Connect(&pg.Options{
		User:     "dev",
		Password: "321321",
		Database: "test",
	})

	old := DEFAULT_TRANSACTION_TIMEOUT
	DEFAULT_TRANSACTION_TIMEOUT = 15 * time.Millisecond

	transaction, err := New(db)
	if !assert.NoError(t, err) {
		t.Fatal("transaction had not been created")
	}

	transaction.Start()

	time.Sleep(20 * time.Millisecond)
	assert.NoError(t, transaction.Commit())

	DEFAULT_TRANSACTION_TIMEOUT = old
}

func TestCommit(t *testing.T) {
	db := pg.Connect(&pg.Options{
		User:     "dev",
		Password: "321321",
		Database: "test",
	})

	transaction, err := New(db)
	if !assert.NoError(t, err) {
		t.Fatal("transaction had not been created")
	}

	transaction.Start()
	assert.NoError(t, transaction.Commit())
}

func TestRollback(t *testing.T) {
	db := pg.Connect(&pg.Options{
		User:     "dev",
		Password: "321321",
		Database: "test",
	})

	transaction, err := New(db)
	if !assert.NoError(t, err) {
		t.Fatal("transaction had not been created")
	}

	time.Sleep(100 * time.Millisecond)

	transaction.Start()
	assert.NoError(t, transaction.Rollback())
}
