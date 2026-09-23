package server

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) attachDeviceLabels(v []*server.Device) error {
	identifiers := make([]int32, 0, len(v))

	for _, d := range v {
		identifiers = append(identifiers, d.Identifier)
	}

	grouped, e := s.store.GroupedLabels(constant.DeviceAddress, identifiers)

	if e != nil {
		return e
	}

	for _, d := range v {
		if labels, okay := grouped[d.Identifier]; okay {
			d.Labels = new(convert.Labels(labels))
		}
	}

	return nil
}
