package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Node(name string) (*proxmox.Node, error) {
	if name == "" {
		return nil, validation.New("node name is required")
	}

	result, e := c.client.Node(c.context, name)

	if e != nil {
		return nil, notFound(e, constant.NodeKind, name)
	}

	return result, nil
}
