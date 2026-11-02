package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListStorages(
	instance string,
	node string,
) (proxmox.Storages, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (proxmox.Storages, error) {
			n, e := c.Node(node)

			if e != nil {
				return nil, e
			}

			return c.Storages(n)
		},
	)
}
