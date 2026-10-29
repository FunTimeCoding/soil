package prometheus

import "context"

func (c *Client) Reachable(x context.Context) error {
	_, e := c.client.Runtimeinfo(x)

	return e
}
