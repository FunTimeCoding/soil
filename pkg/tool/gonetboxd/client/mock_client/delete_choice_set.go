package mock_client

import "context"

func (c *Client) DeleteChoiceSet(_ context.Context, _ string) error {
	return c.Failure
}
