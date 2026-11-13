package service

import (
	"context"
	netboxConstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) Devices(query string) (
	[]*device.Device,
	map[int32][]*label.Label,
	error,
) {
	result, e := peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*device.Device, error) {
			if query != "" {
				return s.client.DevicesByMatch(query)
			}

			return s.client.Devices()
		},
	)

	if e != nil {
		return nil, nil, e
	}

	identifiers := make([]int32, 0, len(result))

	for _, d := range result {
		identifiers = append(identifiers, d.Identifier)
	}

	grouped, e := s.store.GroupedLabels(
		netboxConstant.DeviceAddress,
		identifiers,
	)

	if e != nil {
		return nil, nil, e
	}

	return result, grouped, nil
}
