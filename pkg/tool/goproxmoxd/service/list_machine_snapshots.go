package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) ListMachineSnapshots(
	c face.ProxmoxClient,
	identifier int,
	node string,
) ([]*proxmox.VirtualMachineSnapshot, error) {
	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return nil, e
	}

	return c.MachineSnapshots(vm)
}
