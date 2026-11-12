package mock_client

import "context"

func (c *Client) PostIndexPrune(_ context.Context) ([]string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Removed, nil
}
