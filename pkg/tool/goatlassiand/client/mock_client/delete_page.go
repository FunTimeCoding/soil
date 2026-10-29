package mock_client

import "context"

func (c *Client) DeletePage(
	_ context.Context,
	_ string,
	_ bool,
) error {
	return c.Failure
}
