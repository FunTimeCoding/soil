package mock_client

import "context"

func (c *Client) RefreshSocket(context.Context) error {
	return c.Failure
}
