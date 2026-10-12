package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) Devices(query string) (
	[]*device.Device,
	map[int32][]*label.Label,
	error,
) {
	var result []*device.Device
	var e error

	if query != "" {
		result, e = s.client.DevicesByMatch(query)
	} else {
		result, e = s.client.Devices()
	}

	if e != nil {
		return nil, nil, e
	}

	identifiers := make([]int32, 0, len(result))

	for _, d := range result {
		identifiers = append(identifiers, d.Identifier)
	}

	grouped, e := s.store.GroupedLabels(constant.DeviceAddress, identifiers)

	if e != nil {
		return nil, nil, e
	}

	return result, grouped, nil
}
