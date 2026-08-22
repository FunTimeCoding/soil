package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Node(name string) (*proxmox.Node, error) {
	if name == "" {
		return nil, validation.New("node name is required")
	}

	return c.client.Node(c.context, name)
}
