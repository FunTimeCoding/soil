package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) User(name string) (*proxmox.User, error) {
	if name == "" {
		return nil, validation.New("user name is required")
	}

	return c.client.User(c.context, name)
}
