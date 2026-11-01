package mock_client

import "context"

func (c *Client) PostRun(
	_ context.Context,
	intent string,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	c.Intents = append(c.Intents, intent)

	return nil
}
