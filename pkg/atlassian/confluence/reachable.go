package confluence

import "context"

func (c *Client) Reachable(_ context.Context) error {
	_, e := c.User()

	return e
}
