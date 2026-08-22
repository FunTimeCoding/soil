package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Storage(
	n *proxmox.Node,
	name string,
) (*proxmox.Storage, error) {
	if name == "" {
		return nil, validation.New("storage name is required")
	}

	return n.Storage(c.context, name)
}
