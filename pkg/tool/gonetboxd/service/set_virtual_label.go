package service

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) SetVirtualLabel(
	machine string,
	key string,
	value string,
) (*label.Label, error) {
	m, e := s.virtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	return s.store.SetLabel(
		netboxConstant.VirtualMachineAddress,
		m.Identifier,
		key,
		value,
	)
}
