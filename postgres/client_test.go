package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/pg.v5"
	"gopkg.in/pg.v5/orm"

	clients ".."
)

func TestNew(t *testing.T) {
	db := pg.Connect(&pg.Options{
		User: "dev",
	})

	client := New("test", db)

	assert.True(t, client.Database == db)
}

type fakeCommand struct {
	Message string
}

func (c *fakeCommand) HitDatabase(db orm.DB, fallback clients.Client) error {
	c.Message = "FOOBAR"
	return nil
}

func TestClient_Perform(t *testing.T) {
	db := pg.Connect(&pg.Options{
		User: "dev",
	})

	client := New("test", db)
	command := &fakeCommand{}

	assert.NoError(t, client.Perform(command))
	assert.Equal(t, "FOOBAR", command.Message)
}
