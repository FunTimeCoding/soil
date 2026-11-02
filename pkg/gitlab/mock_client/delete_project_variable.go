package mock_client

import "context"

func (c *Client) DeleteProjectVariable(
	_ context.Context,
	_ int64,
	_ string,
) error {
	return nil
}
