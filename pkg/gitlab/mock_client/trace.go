package mock_client

import "context"

func (c *Client) Trace(
	_ context.Context,
	_ int64,
	job int64,
) (string, error) {
	return c.traces[job], nil
}
