package mock_client

import "context"

func (c *Client) FillNode(
	_ context.Context,
	_ int64,
	_ string,
	_ bool,
) error {
	return nil
}
