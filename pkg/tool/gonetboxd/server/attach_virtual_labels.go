package server

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) attachVirtualLabels(v []*server.VirtualMachine) error {
	identifiers := make([]int32, 0, len(v))

	for _, m := range v {
		identifiers = append(identifiers, m.Identifier)
	}

	grouped, e := s.store.GroupedLabels(
		constant.VirtualMachineAddress,
		identifiers,
	)

	if e != nil {
		return e
	}

	for _, m := range v {
		if labels, okay := grouped[m.Identifier]; okay {
			m.Labels = new(convert.Labels(labels))
		}
	}

	return nil
}
