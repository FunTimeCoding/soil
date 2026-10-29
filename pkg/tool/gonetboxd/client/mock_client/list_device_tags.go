package mock_client

import "context"

func (c *Client) ListDeviceTags(
	_ context.Context,
	_ string,
) ([]string, error) {
	return nil, c.Failure
}
