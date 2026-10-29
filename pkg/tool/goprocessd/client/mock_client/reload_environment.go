package mock_client

import "context"

func (c *Client) ReloadEnvironment(context.Context) error {
	if c.Failure != nil {
		return c.Failure
	}

	c.Reloads++

	return nil
}
