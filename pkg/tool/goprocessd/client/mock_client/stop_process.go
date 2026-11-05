package mock_client

import "context"

func (c *Client) StopProcess(
	_ context.Context,
	name string,
) error {
	return c.setRunning(name, false)
}
