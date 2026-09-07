package mock_client

import "context"

func (c *Client) Navigate(
	_ context.Context,
	_ string,
) error {
	return nil
}
