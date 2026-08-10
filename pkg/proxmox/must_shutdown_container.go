package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustShutdownContainer(v *proxmox.Container) *proxmox.Task {
	result, e := c.ShutdownContainer(v)
	errors.PanicOnError(e)

	return result
}
