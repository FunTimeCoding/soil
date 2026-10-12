package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) VirtualMachines() (
	[]*virtual_machine.Machine,
	map[int32][]*label.Label,
	error,
) {
	result, e := s.client.VirtualMachines()

	if e != nil {
		return nil, nil, e
	}

	identifiers := make([]int32, 0, len(result))

	for _, m := range result {
		identifiers = append(identifiers, m.Identifier)
	}

	grouped, e := s.store.GroupedLabels(
		constant.VirtualMachineAddress,
		identifiers,
	)

	if e != nil {
		return nil, nil, e
	}

	return result, grouped, nil
}
