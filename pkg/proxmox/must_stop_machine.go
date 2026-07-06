package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustStopMachine(v *proxmox.VirtualMachine) *proxmox.Task {
	result, e := c.StopMachine(v)
	errors.PanicOnError(e)

	return result
}
