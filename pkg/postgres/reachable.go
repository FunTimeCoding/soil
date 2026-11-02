package postgres

import "context"

func (c *Client) Reachable(x context.Context) error {
	return c.pool.Ping(x)
}
