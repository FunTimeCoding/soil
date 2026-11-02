package mock_client

import "context"

func (c *Client) CompareFiles(
	_ context.Context,
	_ int64,
	_ string,
	_ string,
) ([]string, error) {
	return nil, nil
}
