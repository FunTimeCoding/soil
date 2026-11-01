package mock_client

import "context"

func (c *Client) DeleteSession(_ context.Context, _ string) error {
	return c.Failure
}
