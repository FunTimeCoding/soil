package mock_client

import "context"

func (c *Client) Restart(
	_ context.Context,
	name string,
) error {
	return c.setRunning(name, true)
}
