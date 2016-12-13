package clients

// Passthrough pass command to underlying client
type Passthrough struct {
	fallback Client
}

// Perform see Client
func (client Passthrough) Perform(command Command) error {
	if client.fallback == nil {
		return NoFallback{}
	}

	return client.fallback.Perform(command)
}

// Wrap see Client
func (client Passthrough) Wrap(fallback Client) Client {
	return Passthrough{
		fallback: fallback,
	}
}
