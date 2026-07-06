package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) GetContainer(
	c face.ProxmoxClient,
	identifier int,
	node string,
) (*proxmox.Container, error) {
	return findContainer(c, identifier, node)
}
