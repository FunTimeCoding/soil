package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) AddNetwork(
	node string,
	iface string,
	bridge string,
) {
	c.networks[node] = append(
		c.networks[node],
		&proxmox.NodeNetwork{Iface: iface, Type: bridge},
	)
}
