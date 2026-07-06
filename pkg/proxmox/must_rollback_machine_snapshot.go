package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustRollbackMachineSnapshot(
	v *proxmox.VirtualMachine,
	name string,
) *proxmox.Task {
	result, e := c.RollbackMachineSnapshot(v, name)
	errors.PanicOnError(e)

	return result
}
