package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustUser(name string) *proxmox.User {
	result, e := c.User(name)
	errors.PanicOnError(e)

	return result
}
