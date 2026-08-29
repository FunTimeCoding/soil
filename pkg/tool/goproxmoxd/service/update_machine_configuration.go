package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) updateMachineConfiguration(
	c face.ProxmoxClient,
	identifier int,
	node string,
	options []proxmox.VirtualMachineOption,
) error {
	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return e
	}

	return c.UpdateMachineConfiguration(vm, options...)
}
