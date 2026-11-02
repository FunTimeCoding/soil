package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListNetworks(
	instance string,
	node string,
) (proxmox.NodeNetworks, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (proxmox.NodeNetworks, error) {
			n, e := c.Node(node)

			if e != nil {
				return nil, e
			}

			return c.Networks(n)
		},
	)
}
