package client

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func (c *Client) PulseCruise(
	sessionIdentifier string,
	body string,
) (bool, error) {
	return c.pulseWith(sessionIdentifier, body, client.Cruise)
}
