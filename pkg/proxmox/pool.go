package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Pool(name string) (*proxmox.Pool, error) {
	if name == "" {
		return nil, validation.New("pool name is required")
	}

	return c.client.Pool(c.context, name)
}
