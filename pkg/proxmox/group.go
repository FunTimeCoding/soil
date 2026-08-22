package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Group(name string) (*proxmox.Group, error) {
	if name == "" {
		return nil, validation.New("group name is required")
	}

	return c.client.Group(c.context, name)
}
