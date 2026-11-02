package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) Networks(n *proxmox.Node) (proxmox.NodeNetworks, error) {
	return c.networks[n.Name], nil
}
