package band

import "context"

func (c *Client) Reachable(_ context.Context) error {
	_, e := c.GeneralSettings()

	return e
}
