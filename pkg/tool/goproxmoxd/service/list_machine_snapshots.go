package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListMachineSnapshots(
	instance string,
	identifier int,
	node string,
) ([]*proxmox.VirtualMachineSnapshot, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) ([]*proxmox.VirtualMachineSnapshot, error) {
			vm, e := findMachine(c, identifier, node)

			if e != nil {
				return nil, e
			}

			return c.MachineSnapshots(vm)
		},
	)
}
