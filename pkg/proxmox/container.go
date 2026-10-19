package proxmox

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
)

func (c *Client) Container(
	n *proxmox.Node,
	identifier int,
) (*proxmox.Container, error) {
	result, e := n.Container(c.context, identifier)

	if e != nil {
		return nil, notFound(e, constant.ContainerKind, identifier)
	}

	return result, nil
}
