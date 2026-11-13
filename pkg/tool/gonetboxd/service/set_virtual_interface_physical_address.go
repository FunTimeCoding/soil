package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/netbox-community/go-netbox/v4"
)

func (s *Service) SetVirtualInterfacePhysicalAddress(
	machine string,
	interfaceName string,
	address string,
) (*netbox.VMInterface, error) {
	h, e := physicalAddress(address)

	if e != nil {
		return nil, e
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*netbox.VMInterface, error) {
			m, f := s.client.VirtualMachine(machine)

			if f != nil {
				return nil, f
			}

			return s.client.UpdateVirtualInterface(m, interfaceName, h)
		},
	)
}
