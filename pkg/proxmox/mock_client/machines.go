package mock_client

import (
	"fmt"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Machines(n *proxmox.Node) (proxmox.VirtualMachines, error) {
	nodeMachines, okay := c.machines[n.Name]

	if !okay {
		return nil, fmt.Errorf("node %q not found", n.Name)
	}

	var result proxmox.VirtualMachines

	for _, vm := range nodeMachines {
		result = append(result, vm)
	}

	return result, nil
}
