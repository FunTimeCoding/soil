package proxmox

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Machine(
	n *proxmox.Node,
	identifier int,
) (*proxmox.VirtualMachine, error) {
	result, e := n.VirtualMachine(c.context, identifier)

	if e != nil {
		return nil, notFound(e, constant.MachineKind, identifier)
	}

	return result, nil
}
