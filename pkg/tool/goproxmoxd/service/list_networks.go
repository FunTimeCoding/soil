package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListNetworks(
	c face.ProxmoxClient,
	node string,
) (proxmox.NodeNetworks, error) {
	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	return c.Networks(n)
}
