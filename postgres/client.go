package postgres

import (
	"gopkg.in/pg.v5"
	"gopkg.in/pg.v5/orm"

	"gopkg.in/gopaws/go.clients.v1"
	"gopkg.in/gopaws/go.clients.v1/postgres/transaction"
)

// New return new PostgreSQL connection wrapper
func New(key clients.Key, database *pg.DB) *Client {
	client := new(Client)
	client.Database = database
	client.key = key
	client.transactions = transaction.NewManager(database)
	client.fallback = clients.Null(nil)

	return client
}

// Client is an implementaion of clients.Client
type Client struct {
	Database *pg.DB

	key          clients.Key
	transactions *transaction.Manager
	fallback     clients.Client
}

// Action
type Action interface {
	HitPostgres(key clients.Key, connection orm.DB, fallback clients.Client) error
}

// Transaction return new database transaction object and error
func (client Client) Transaction(operation string) (clients.Transaction, error) {
	return client.transactions.Start(operation)
}

// Perform see clients.Client
func (client Client) Perform(command clients.Command) error {
	if action, ok := command.(Action); ok {
		if tx, ok := action.(clients.TransactionCommand); ok {
			transaction, err := client.transactions.Checkout(tx.Transaction())
			if err != nil {
				return err
			}

			return action.HitPostgres(client.key, transaction, client.fallback)
		}

		return action.HitPostgres(client.key, client.Database, client.fallback)
	}

	return client.fallback.Perform(command)
}

// Wrap see clients.Client
func (client Client) Wrap(fallback clients.Client) clients.Client {
	if fallback == nil {
		return &client
	}

	return &Client{
		Database:     client.Database,
		key:          client.key,
		transactions: client.transactions,
		fallback:     fallback,
	}
}
