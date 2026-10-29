package mock_client

import "context"

func (c *Client) DeleteCustomField(_ context.Context, _ string) error {
	return c.Failure
}
