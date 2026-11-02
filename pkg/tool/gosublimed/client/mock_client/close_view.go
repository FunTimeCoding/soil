package mock_client

import "context"

func (c *Client) CloseView(
	_ context.Context,
	_ int,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.Failure
}
