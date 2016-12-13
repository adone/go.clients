package clients

import (
	"sync"
)

type Key string

// Registry contains all clients used by application
type Registry struct {
	mutex sync.RWMutex

	clients map[Key]Client
}

// Add stores provided client in regsitry
func (registry *Registry) Add(key Key, client Client) {
	registry.mutex.Lock()
	defer registry.mutex.Unlock()

	registry.clients[key] = client
}

// Get returns client linked with a key
func (registry *Registry) Get(key Key) Client {
	registry.mutex.RLock()
	defer registry.mutex.RUnlock()

	if client, ok := registry.clients[key]; ok {
		return client
	}

	return Passthrough{Null(NoFallback{})}
}

var registry Registry

func init() {
	registry = Registry{
		clients: make(map[Key]Client),
	}
}

// Get returns client linked with a key
func Get(key Key) Client {
	return registry.Get(key)
}

// Add link provided client with a key
func Add(key Key, client Client) {
	registry.Add(key, client)
}
