package mock_client

import "context"

func (c *Client) UpdateComment(
	_ context.Context,
	_ string,
	_ string,
	_ string,
) error {
	return c.Failure
}
