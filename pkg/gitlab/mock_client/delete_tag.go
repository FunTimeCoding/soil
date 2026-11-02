package mock_client

import "context"

func (c *Client) DeleteTag(
	_ context.Context,
	_ int64,
	_ string,
) error {
	return c.failure
}
