package mock_client

import "context"

func (c *Client) Query(
	_ context.Context,
	_ string,
	_ any,
) error {
	return c.failure
}
