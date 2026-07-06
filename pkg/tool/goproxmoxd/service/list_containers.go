package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListContainers(
	c face.ProxmoxClient,
	node string,
) (proxmox.Containers, error) {
	nodeNames, e := s.resolveNodeNames(c, node)

	if e != nil {
		return nil, e
	}

	var result proxmox.Containers

	for _, name := range nodeNames {
		n, e := c.Node(name)

		if e != nil {
			return nil, e
		}

		containers, e := c.Containers(n)

		if e != nil {
			return nil, e
		}

		result = append(result, containers...)
	}

	return result, nil
}
