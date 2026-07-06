package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustPools() []*proxmox.Pool {
	result, e := c.Pools()
	errors.PanicOnError(e)

	return result
}
