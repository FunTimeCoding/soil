package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustVersion() *proxmox.Version {
	result, e := c.Version()
	errors.PanicOnError(e)

	return result
}
