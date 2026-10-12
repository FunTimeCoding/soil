package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) SetVirtualLabel(
	machine string,
	key string,
	value string,
) (*label.Label, error) {
	m, e := s.client.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	return s.store.SetLabel(
		constant.VirtualMachineAddress,
		m.Identifier,
		key,
		value,
	)
}
