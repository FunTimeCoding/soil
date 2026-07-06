package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListContainerSnapshots(
	c face.ProxmoxClient,
	identifier int,
	node string,
) ([]*proxmox.ContainerSnapshot, error) {
	ct, e := findContainer(c, identifier, node)

	if e != nil {
		return nil, e
	}

	return c.ContainerSnapshots(ct)
}
