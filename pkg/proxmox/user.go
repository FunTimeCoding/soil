package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) User(name string) (*proxmox.User, error) {
	if name == "" {
		return nil, validation.New("user name is required")
	}

	result, e := c.client.User(c.context, name)

	if e != nil {
		return nil, notFound(e, constant.UserKind, name)
	}

	return result, nil
}
