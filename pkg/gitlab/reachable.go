package gitlab

import "context"

func (c *Client) Reachable(_ context.Context) error {
	_, e := c.CurrentUser()

	return e
}
