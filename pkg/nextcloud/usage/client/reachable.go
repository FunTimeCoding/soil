package client

import "context"

func (c *Client) Reachable(_ context.Context) error {
	_, e := c.Fetch()

	return e
}
