package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustStartContainer(v *proxmox.Container) *proxmox.Task {
	result, e := c.StartContainer(v)
	errors.PanicOnError(e)

	return result
}
