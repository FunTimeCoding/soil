package mock_client

import "context"

func (c *Client) DeleteConfigurationContext(_ context.Context, _ string) error {
	return c.Failure
}
