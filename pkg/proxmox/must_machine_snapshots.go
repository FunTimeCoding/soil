package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustMachineSnapshots(v *proxmox.VirtualMachine) []*proxmox.VirtualMachineSnapshot {
	result, e := c.MachineSnapshots(v)
	errors.PanicOnError(e)

	return result
}
