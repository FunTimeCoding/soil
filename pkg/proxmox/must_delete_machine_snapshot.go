package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustDeleteMachineSnapshot(
	v *proxmox.VirtualMachine,
	name string,
) *proxmox.Task {
	result, e := c.DeleteMachineSnapshot(v, name)
	errors.PanicOnError(e)

	return result
}
