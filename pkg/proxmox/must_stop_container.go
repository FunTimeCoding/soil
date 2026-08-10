package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustStopContainer(v *proxmox.Container) *proxmox.Task {
	result, e := c.StopContainer(v)
	errors.PanicOnError(e)

	return result
}
