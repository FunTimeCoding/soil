package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) VirtualInterfaces(machine string) ([]*network.Interface, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*network.Interface, error) {
			m, e := s.client.VirtualMachine(machine)

			if e != nil {
				return nil, e
			}

			return s.client.VirtualMachineInterfaces(m)
		},
	)
}
