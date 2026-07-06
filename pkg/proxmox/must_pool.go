package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustPool(name string) *proxmox.Pool {
	result, e := c.Pool(name)
	errors.PanicOnError(e)

	return result
}
