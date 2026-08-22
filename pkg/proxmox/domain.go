package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Domain(name string) (*proxmox.Domain, error) {
	if name == "" {
		return nil, validation.New("domain name is required")
	}

	return c.client.Domain(c.context, name)
}
