package transaction

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/pg.v5"
)

func TestNewManager(t *testing.T) {
	db := pg.Connect(&pg.Options{
		User:     "dev",
		Password: "321321",
		Database: "test",
	})

	manager := NewManager(db)
	if assert.NotNil(t, manager) {
		assert.Equal(t, uint64(0), manager.counter)
	}
}

func TestStart(t *testing.T) {
	db := pg.Connect(&pg.Options{
		User:     "dev",
		Password: "321321",
		Database: "test",
	})

	manager := NewManager(db)

	old := DEFAULT_TRANSACTION_TIMEOUT
	DEFAULT_TRANSACTION_TIMEOUT = 15 * time.Millisecond

	_, err := manager.Start("test")
	if !assert.NoError(t, err) {
		t.Fatal("transaction had not been created")
	}

	time.Sleep(100 * time.Millisecond)

	_, err = manager.Checkout("test-1")
	assert.Error(t, err)

	DEFAULT_TRANSACTION_TIMEOUT = old
}

func TestCheckout(t *testing.T) {
	db := pg.Connect(&pg.Options{
		User:     "dev",
		Password: "321321",
		Database: "test",
	})

	manager := NewManager(db)
	transaction, err := manager.Start("test")
	if !assert.NoError(t, err) {
		t.Fatal("transaction had not been created")
	}

	_, err = manager.Checkout("test-1")
	assert.NoError(t, err)

	assert.NoError(t, transaction.Commit())

	_, err = manager.Checkout("test-1")
	assert.Error(t, err)
}
