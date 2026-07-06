package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustNetworks(n *proxmox.Node) proxmox.NodeNetworks {
	result, e := c.Networks(n)
	errors.PanicOnError(e)

	return result
}
