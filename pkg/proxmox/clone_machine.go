package proxmox

import (
	"github.com/funtimecoding/soil/pkg/proxmox/refusal_transport"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) CloneMachine(
	vm *proxmox.VirtualMachine,
	options *proxmox.VirtualMachineCloneOptions,
) (int, *proxmox.Task, error) {
	return vm.Clone(refusal_transport.WithCaller(c.context), options)
}
