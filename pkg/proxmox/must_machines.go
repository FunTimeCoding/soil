package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) MustMachines(n *proxmox.Node) proxmox.VirtualMachines {
	result, e := c.Machines(n)
	errors.PanicOnError(e)

	return result
}
