package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustNode(name string) *proxmox.Node {
	result, e := c.Node(name)
	errors.PanicOnError(e)

	return result
}
