package service

import "github.com/funtimecoding/soil/pkg/netbox/constant"

func (s *Service) RemoveVirtualLabel(
	machine string,
	key string,
) error {
	m, e := s.client.VirtualMachineByName(machine)

	if e != nil {
		return e
	}

	return s.store.RemoveLabel(
		constant.VirtualMachineAddress,
		m.Identifier,
		key,
	)
}
