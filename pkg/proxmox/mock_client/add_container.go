package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) AddContainer(
	node string,
	identifier int,
	name string,
) {
	c.containers[node][identifier] = &proxmox.Container{
		VMID: proxmox.StringOrUint64(identifier),
		Name: name,
		Node: node,
	}
}
