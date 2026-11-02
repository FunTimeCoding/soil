package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListContainers(
	instance string,
	node string,
) (proxmox.Containers, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (proxmox.Containers, error) {
			nodeNames, e := s.resolveNodeNames(c, node)

			if e != nil {
				return nil, e
			}

			var result proxmox.Containers

			for _, name := range nodeNames {
				n, f := c.Node(name)

				if f != nil {
					return nil, f
				}

				containers, f := c.Containers(n)

				if f != nil {
					return nil, f
				}

				result = append(result, containers...)
			}

			return result, nil
		},
	)
}
