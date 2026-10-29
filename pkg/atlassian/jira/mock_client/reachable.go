package mock_client

import "context"

func (c *Client) Reachable(_ context.Context) error {
	return nil
}
