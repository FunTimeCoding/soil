package service

import (
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) VirtualLabels(machine string) ([]*label.Label, error) {
	m, e := s.client.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	return s.store.Labels(netboxConstant.VirtualMachineAddress, m.Identifier)
}
