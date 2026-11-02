package proxmox

import (
	"github.com/funtimecoding/soil/pkg/proxmox/refusal_transport"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) CreateMachine(
	n *proxmox.Node,
	identifier int,
	options ...proxmox.VirtualMachineOption,
) (*proxmox.Task, error) {
	return n.NewVirtualMachine(
		refusal_transport.WithCaller(c.context),
		identifier,
		options...,
	)
}
