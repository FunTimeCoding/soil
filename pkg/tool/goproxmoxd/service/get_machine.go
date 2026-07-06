package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) GetMachine(
	c face.ProxmoxClient,
	identifier int,
	node string,
) (*proxmox.VirtualMachine, error) {
	return findMachine(c, identifier, node)
}
