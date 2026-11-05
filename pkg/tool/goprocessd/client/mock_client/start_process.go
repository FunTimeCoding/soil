package mock_client

import "context"

func (c *Client) StartProcess(
	_ context.Context,
	name string,
) error {
	return c.setRunning(name, true)
}
