package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustResetMachine(v *proxmox.VirtualMachine) *proxmox.Task {
	result, e := c.ResetMachine(v)
	errors.PanicOnError(e)

	return result
}
