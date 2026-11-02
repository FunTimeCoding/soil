package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) GetMachine(
	instance string,
	identifier int,
	node string,
) (*proxmox.VirtualMachine, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (*proxmox.VirtualMachine, error) {
			return findMachine(c, identifier, node)
		},
	)
}
