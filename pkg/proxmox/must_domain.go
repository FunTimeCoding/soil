package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustDomain(name string) *proxmox.Domain {
	result, e := c.Domain(name)
	errors.PanicOnError(e)

	return result
}
