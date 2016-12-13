package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullClient_Perform(t *testing.T) {
	client := Null(nil)
	assert.NoError(t, client.Perform("Foobar"))

	clientError := Null(NoFallback{})
	assert.Error(t, clientError.Perform("Foobar"))
}

func TestNullClient_Wrap(t *testing.T) {
	client := Null(nil)

	assert.Equal(t, client, client.Wrap(Passthrough{}))
}
