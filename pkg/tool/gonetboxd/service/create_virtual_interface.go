package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/netbox-community/go-netbox/v4"
)

func (s *Service) CreateVirtualInterface(
	machine string,
	name string,
) (*netbox.VMInterface, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*netbox.VMInterface, error) {
			m, e := s.client.VirtualMachine(machine)

			if e != nil {
				return nil, e
			}

			return s.client.CreateVirtualInterface(m, name)
		},
	)
}
