package connector

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func (c *Client) Pulse(
	sessionIdentifier string,
	body string,
) (bool, error) {
	return c.pulseWith(sessionIdentifier, body, client.Queued)
}
