package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustUsers() []*proxmox.User {
	result, e := c.Users()
	errors.PanicOnError(e)

	return result
}
