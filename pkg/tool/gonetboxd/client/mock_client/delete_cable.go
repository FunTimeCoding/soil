package mock_client

import "context"

func (c *Client) DeleteCable(
	_ context.Context,
	_ int32,
) error {
	return c.Failure
}
