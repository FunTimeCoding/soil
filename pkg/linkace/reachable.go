package linkace

import "context"

func (c *Client) Reachable(_ context.Context) error {
	_, e := c.ListsPage(1)

	return e
}
