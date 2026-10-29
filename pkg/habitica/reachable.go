package habitica

import "context"

func (c *Client) Reachable(_ context.Context) error {
	_, e := c.user()

	return e
}
