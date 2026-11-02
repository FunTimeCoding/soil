package mock_client

import "context"

func (c *Client) DeleteBranch(
	_ context.Context,
	_ int32,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Failure
}
