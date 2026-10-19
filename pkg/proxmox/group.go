package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Group(name string) (*proxmox.Group, error) {
	if name == "" {
		return nil, validation.New("group name is required")
	}

	result, e := c.client.Group(c.context, name)

	if e != nil {
		return nil, notFound(e, constant.GroupKind, name)
	}

	return result, nil
}
