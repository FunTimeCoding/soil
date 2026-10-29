package mock_client

import "context"

func (c *Client) ReloadProcfile(context.Context) error {
	if c.Failure != nil {
		return c.Failure
	}

	c.Reloads++

	return nil
}
