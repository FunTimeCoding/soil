package service

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (s *Service) SetVirtualMachineCustomFields(
	name string,
	fields map[string]any,
) (*virtual_machine.Machine, error) {
	return s.client.SetVirtualMachineCustomFields(name, fields)
}
