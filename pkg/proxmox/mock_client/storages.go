package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Storages(n *proxmox.Node) (proxmox.Storages, error) {
	return c.storages[n.Name], nil
}
