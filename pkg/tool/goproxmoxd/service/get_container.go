package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) GetContainer(
	instance string,
	identifier int,
	node string,
) (*proxmox.Container, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (*proxmox.Container, error) {
			return findContainer(c, identifier, node)
		},
	)
}
