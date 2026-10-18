package ollama

import "context"

func (c *Client) Reachable(x context.Context) error {
	return c.client.Heartbeat(x)
}
