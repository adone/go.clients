package redis

import (
	"gopkg.in/gopaws/go.clients.v1"
	"gopkg.in/gopaws/go.redis.v1/storage"
)

// New return new redis storage wrapper
func New(key clients.Key, redis *storage.Client) *Client {
	client := new(Client)
	client.key = key
	client.storage = redis
	client.fallback = clients.Null(nil)

	return client
}

// Client is an implementation of clients.Client
type Client struct {
	key      clients.Key
	storage  *storage.Client
	fallback clients.Client
}

// Action
type Action interface {
	HitRedis(clients.Key, *storage.Client, clients.Client) error
}

// Perform see clients.Client
func (client Client) Perform(command clients.Command) error {
	if action, ok := command.(Action); ok {
		return action.HitRedis(client.key, client.storage, client.fallback)
	}

	return client.fallback.Perform(command)
}

// Wrap see clients.Client
func (client Client) Wrap(fallback clients.Client) clients.Client {
	if fallback == nil {
		return &client
	}

	return &Client{
		key:      client.key,
		storage:  client.storage,
		fallback: fallback,
	}
}
