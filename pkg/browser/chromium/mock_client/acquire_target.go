package mock_client

import "context"

func (c *Client) AcquireTarget(_ string) (context.Context, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return context.Background(), nil
}
