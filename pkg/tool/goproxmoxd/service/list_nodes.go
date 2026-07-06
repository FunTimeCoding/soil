package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListNodes(
	c face.ProxmoxClient,
) (proxmox.NodeStatuses, error) {
	return c.Nodes()
}
