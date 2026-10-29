package mock_client

import "context"

func (c *Client) ClearLog(
	_ context.Context,
	name string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	if _, e := c.index(name); e != nil {
		return e
	}

	delete(c.Logs, name)

	return nil
}
