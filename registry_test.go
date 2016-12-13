package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistry_Add(t *testing.T) {
	var a Key = "foobar"
	client := NullClient{}

	registry.Add(a, client)
	assert.True(t, client == registry.clients[a])
}

func TestRegistry_Get(t *testing.T) {
	var a Key = "foobar"
	client := NullClient{}

	registry.Add(a, client)
	assert.True(t, client == registry.Get(a))
}

func TestAdd(t *testing.T) {
	var a Key = "foobar"
	client := NullClient{}

	Add(a, client)
	assert.True(t, client == registry.clients[a])
}

func TestGet(t *testing.T) {
	var a Key = "foobar"
	client := NullClient{}

	Add(a, client)
	assert.True(t, client == Get(a))
}
