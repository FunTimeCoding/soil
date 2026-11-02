package proxmox

import (
	"github.com/funtimecoding/soil/pkg/proxmox/refusal_transport"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) UpdateMachineConfiguration(
	v *proxmox.VirtualMachine,
	options ...proxmox.VirtualMachineOption,
) error {
	return v.ConfigSync(refusal_transport.WithCaller(c.context), options...)
}
