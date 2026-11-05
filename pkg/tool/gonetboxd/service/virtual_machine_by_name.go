package service

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (s *Service) virtualMachineByName(
	name string,
) (*virtual_machine.Machine, error) {
	return guarded(
		s.guard,
		func() (*virtual_machine.Machine, error) {
			return s.client.VirtualMachineByName(name)
		},
	)
}
