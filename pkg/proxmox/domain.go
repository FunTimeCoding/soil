package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Domain(name string) (*proxmox.Domain, error) {
	if name == "" {
		return nil, validation.New("domain name is required")
	}

	result, e := c.client.Domain(c.context, name)

	if e != nil {
		return nil, notFound(e, constant.DomainKind, name)
	}

	return result, nil
}
