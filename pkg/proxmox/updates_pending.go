package proxmox

import "github.com/luthermonson/go-proxmox"

func (c *Client) UpdatesPending(n *proxmox.Node) ([]*proxmox.APTUpdate, error) {
	return n.APTUpdates(c.context)
}
