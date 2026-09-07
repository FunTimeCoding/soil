package mock_client

import "context"

func (c *Client) Evaluate(
	_ context.Context,
	_ string,
	_ any,
) error {
	return nil
}
