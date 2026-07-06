package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustRole(name string) *proxmox.Permission {
	result, e := c.Role(name)
	errors.PanicOnError(e)

	return result
}
