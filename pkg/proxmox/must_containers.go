package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustContainers(n *proxmox.Node) proxmox.Containers {
	result, e := c.Containers(n)
	errors.PanicOnError(e)

	return result
}
