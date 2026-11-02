package proxmox

import "context"

func (c *Client) Reachable(x context.Context) error {
	_, e := c.client.Version(x)

	return e
}
