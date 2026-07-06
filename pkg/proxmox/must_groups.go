package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustGroups() proxmox.Groups {
	result, e := c.Groups()
	errors.PanicOnError(e)

	return result
}
