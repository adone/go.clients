package clients

// Client
type Client interface {
	Perform(Command) error
	Wrap(Client) Client
}

// Command
type Command interface{}

// TransactionCommand
type TransactionCommand interface {
	Transaction() string
}

// Transaction
type Transaction interface {
	Key() string
	Rollback() error
	Commit() error
}
