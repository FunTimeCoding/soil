package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Machine(
	n *proxmox.Node,
	identifier int,
) (*proxmox.VirtualMachine, error) {
	vm, okay := c.machines[n.Name][identifier]

	if !okay {
		return nil, not_found.New(constant.MachineKind, identifier)
	}

	return vm, nil
}
