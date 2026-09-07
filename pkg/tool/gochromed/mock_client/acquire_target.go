package mock_client

import "context"

func (c *Client) AcquireTarget(_ string) context.Context {
	return context.Background()
}
