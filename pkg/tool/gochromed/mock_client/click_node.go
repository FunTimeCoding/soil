package mock_client

import "context"

func (c *Client) ClickNode(
	_ context.Context,
	_ int64,
) error {
	return nil
}
