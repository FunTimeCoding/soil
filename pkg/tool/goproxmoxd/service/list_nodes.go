package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListNodes(instance string) (proxmox.NodeStatuses, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (proxmox.NodeStatuses, error) {
			return c.Nodes()
		},
	)
}
