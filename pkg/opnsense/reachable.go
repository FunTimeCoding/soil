package opnsense

import "context"

func (c *Client) Reachable(context.Context) error {
	_, e := c.Interfaces()

	return e
}
