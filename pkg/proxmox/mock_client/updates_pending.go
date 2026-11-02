package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) UpdatesPending(n *proxmox.Node) ([]*proxmox.APTUpdate, error) {
	return c.updatesPending[n.Name], nil
}
