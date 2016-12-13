/*
Package mysql Клиент для работы с БД.
*/
package mysql

import (
	"github.com/jinzhu/gorm"

	"gopkg.in/adone/go.clients.v1"
	"gopkg.in/adone/go.clients.v1/mysql/transaction"
)

// New create new MySQL client wrapper
func New(key clients.Key, connection *gorm.DB) *Client {
	client := new(Client)
	client.key = key
	client.connection = connection
	client.transactions = transaction.NewManager(connection)
	client.fallback = clients.Null(nil)

	return client
}

// Client is an implementaion of clients.Client
type Client struct {
	key          clients.Key
	transactions *transaction.Manager
	connection   *gorm.DB
	fallback     clients.Client
}

// Action
type Action interface {
	HitMysql(key clients.Key, connection *gorm.DB, fallback clients.Client) error
}

// Transaction returns new database transaction and error
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

			return action.HitMysql(client.key, transaction, client.fallback)
		}

		return action.HitMysql(client.key, client.connection, client.fallback)
	}

	if client.fallback == nil {
		return clients.NoFallback{}
	}

	return client.fallback.Perform(command)
}

// Wrap see clients.Client
func (client Client) Wrap(fallback clients.Client) clients.Client {
	if fallback == nil {
		return &client
	}

	return &Client{
		key:          client.key,
		transactions: client.transactions,
		connection:   client.connection,
		fallback:     fallback,
	}
}
