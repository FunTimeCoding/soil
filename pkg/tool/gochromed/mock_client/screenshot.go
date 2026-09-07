package mock_client

import "context"

func (c *Client) Screenshot(_ context.Context) ([]byte, error) {
	return nil, nil
}
