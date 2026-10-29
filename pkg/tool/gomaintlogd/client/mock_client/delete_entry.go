package mock_client

import "context"

func (c *Client) DeleteEntry(
	_ context.Context,
	_ int,
) error {
	return c.Failure
}
