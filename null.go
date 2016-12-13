package clients

// NoFallback error
type NoFallback struct{}

// Error interface
func (fallback NoFallback) Error() string {
	return "No fallback"
}

// Null returns a NullClient
func Null(err error) NullClient {
	if err != nil {
		return NullClient{err}
	}

	return NullClient{}
}

// NullClient provides an empty client for chain terminaion
// it contains optional error for result of Perform call
type NullClient struct {
	Error error
}

// Perform see Client
func (client NullClient) Perform(command Command) error {
	return client.Error
}

// Wrap see Client
func (client NullClient) Wrap(Client) Client {
	return client
}
