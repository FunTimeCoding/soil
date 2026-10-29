package mock_client

import "context"

func (c *Client) DeleteLink(
	_ context.Context,
	_ string,
) error {
	return c.Failure
}
