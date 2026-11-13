package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateVirtualMachine(
	name string,
	cluster string,
) (*virtual_machine.Machine, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*virtual_machine.Machine, error) {
			c, e := s.client.Cluster(cluster)

			if e != nil {
				return nil, e
			}

			return s.client.CreateVirtualMachine(name, c)
		},
	)
}
