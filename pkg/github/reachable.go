package github

import "context"

func (c *Client) Reachable(x context.Context) error {
	_, _, e := c.client.RateLimit.Get(x)

	return e
}
