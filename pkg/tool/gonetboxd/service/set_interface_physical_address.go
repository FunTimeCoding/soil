package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SetInterfacePhysicalAddress(
	device string,
	interfaceName string,
	address string,
) (*network.Interface, error) {
	h, e := physicalAddress(address)

	if e != nil {
		return nil, e
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*network.Interface, error) {
			d, f := s.client.Device(device)

			if f != nil {
				return nil, f
			}

			i, f := s.client.DeviceInterface(d, interfaceName)

			if f != nil {
				return nil, f
			}

			return s.client.UpdateInterface(d, interfaceName, i.Type, h)
		},
	)
}
