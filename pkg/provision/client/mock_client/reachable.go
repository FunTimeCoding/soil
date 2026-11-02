package mock_client

import "context"

func (c *Client) Reachable(context.Context) error {
	return c.failures.Caller()
}
