package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustGroup(name string) *proxmox.Group {
	result, e := c.Group(name)
	errors.PanicOnError(e)

	return result
}
