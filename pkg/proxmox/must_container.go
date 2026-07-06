package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustContainer(
	n *proxmox.Node,
	identifier int,
) *proxmox.Container {
	result, e := c.Container(n, identifier)
	errors.PanicOnError(e)

	return result
}
