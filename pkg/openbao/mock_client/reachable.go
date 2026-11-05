package mock_client

import "context"

func (c *Client) Reachable(context.Context) error {
	if c.Failure != nil {
		return c.Failure
	}

	return c.health
}
