package mock_client

import "context"

func (c *Client) ReconfigureDnsmasq(context.Context) error {
	if c.Failure != nil {
		return c.Failure
	}

	c.Reconfigured++

	return nil
}
