package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustNodes() proxmox.NodeStatuses {
	result, e := c.Nodes()
	errors.PanicOnError(e)

	return result
}
