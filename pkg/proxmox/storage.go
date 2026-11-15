package proxmox

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Storage(
	n *proxmox.Node,
	name string,
) (*proxmox.Storage, error) {
	if name == "" {
		return nil, validation.New("storage name is required")
	}

	result := n.StorageRef(name)

	if e := result.Ping(c.context); e != nil {
		return nil, notFound(e, constant.StorageKind, name)
	}

	return result, nil
}
