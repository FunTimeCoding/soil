package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListContainerSnapshots(
	instance string,
	identifier int,
	node string,
) ([]*proxmox.ContainerSnapshot, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) ([]*proxmox.ContainerSnapshot, error) {
			ct, e := findContainer(c, identifier, node)

			if e != nil {
				return nil, e
			}

			return c.ContainerSnapshots(ct)
		},
	)
}
