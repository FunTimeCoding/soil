package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Role(name string) (*proxmox.Permission, error) {
	if name == "" {
		return nil, validation.New("role name is required")
	}

	result, e := c.client.Role(c.context, name)

	if e != nil {
		return nil, notFound(e, constant.RoleKind, name)
	}

	return &result, nil
}
