package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func findContainer(
	c face.ProxmoxClient,
	identifier int,
	nodeName string,
) (*proxmox.Container, error) {
	if nodeName != "" {
		node, e := c.Node(nodeName)

		if e != nil {
			return nil, e
		}

		return c.Container(node, identifier)
	}

	nodes, e := c.Nodes()

	if e != nil {
		return nil, e
	}

	for _, n := range nodes {
		node, f := c.Node(n.Node)

		if f != nil {
			continue
		}

		containers, g := c.Containers(node)

		if g != nil {
			continue
		}

		for _, listed := range containers {
			if int(listed.VMID) == identifier {
				return c.Container(node, identifier)
			}
		}
	}

	return nil, not_found.New("container", identifier)
}
