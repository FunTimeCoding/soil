package sentry

import "context"

func (c *Client) Reachable(_ context.Context) error {
	_, e := c.Whoami()

	return e
}
