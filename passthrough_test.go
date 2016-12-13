package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassthrough_Perform(t *testing.T) {
	client := Passthrough{
		fallback: Passthrough{},
	}

	assert.Error(t, client.Perform("Foobar"))
	assert.Equal(t, NoFallback{}, client.Perform("Foobar"))
}
