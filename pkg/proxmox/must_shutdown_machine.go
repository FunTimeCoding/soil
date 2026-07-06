package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustShutdownMachine(v *proxmox.VirtualMachine) *proxmox.Task {
	result, e := c.ShutdownMachine(v)
	errors.PanicOnError(e)

	return result
}
