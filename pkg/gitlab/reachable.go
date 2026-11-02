package gitlab

import "context"

func (c *Client) Reachable(x context.Context) error {
	_, e := c.CurrentUser(x)

	return e
}
