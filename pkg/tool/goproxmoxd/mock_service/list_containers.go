package mock_service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListContainers(
	c face.ProxmoxClient,
	node string,
) (proxmox.Containers, error) {
	if node == "" {
		return nil, nil
	}

	n, e := c.Node(node)

	if e != nil {
		return nil, e
	}

	return c.Containers(n)
}
