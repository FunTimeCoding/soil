package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Container(
	n *proxmox.Node,
	identifier int,
) (*proxmox.Container, error) {
	ct, okay := c.containers[n.Name][identifier]

	if !okay {
		return nil, not_found.New(constant.ContainerKind, identifier)
	}

	return ct, nil
}
