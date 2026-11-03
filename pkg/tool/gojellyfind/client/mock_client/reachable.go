package mock_client

import "context"

func (c *Client) Reachable(_ context.Context) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	return c.failures.Caller()
}
