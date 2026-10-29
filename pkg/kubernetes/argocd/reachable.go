package argocd

import "context"

func (c *Client) Reachable(_ context.Context) error {
	_, e := c.Applications()

	return e
}
