package transaction

import (
	"fmt"
)

type NotFound struct {
	key string
}

func (err NotFound) Error() string {
	return fmt.Sprintf("Transaction with key %s not found", err.key)
}
