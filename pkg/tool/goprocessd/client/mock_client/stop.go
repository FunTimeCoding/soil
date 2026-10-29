package mock_client

import "context"

func (c *Client) Stop(
	_ context.Context,
	name string,
) error {
	return c.setRunning(name, false)
}
