package service

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) VirtualMachine(name string) (
	*virtual_machine.Machine,
	[]*label.Label,
	error,
) {
	m, e := s.virtualMachineByName(name)

	if e != nil {
		return nil, nil, e
	}

	labels, e := s.store.Labels(
		netboxConstant.VirtualMachineAddress,
		m.Identifier,
	)

	if e != nil {
		return nil, nil, e
	}

	return m, labels, nil
}
