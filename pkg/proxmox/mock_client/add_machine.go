package mock_client

import "github.com/luthermonson/go-proxmox"

func (c *Client) AddMachine(
	node string,
	identifier int,
	name string,
) {
	c.machines[node][identifier] = &proxmox.VirtualMachine{
		VMID: proxmox.StringOrUint64(identifier),
		Name: name,
		Node: node,
	}
}
