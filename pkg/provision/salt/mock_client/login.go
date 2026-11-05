package mock_client

import "context"

func (c *Client) Login(context.Context) error {
	return c.failures.Caller()
}
