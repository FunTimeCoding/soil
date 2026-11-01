package mock_client

import "context"

func (c *Client) Reachable(context.Context) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Probe
}
