package mock_client

import "context"

func (c *Client) DeletePackage(
	_ context.Context,
	_ int64,
	_ int64,
) error {
	return c.failure
}
