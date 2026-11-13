package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) VirtualDisks(machine string) ([]*virtual_disk.Disk, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*virtual_disk.Disk, error) {
			m, e := s.client.VirtualMachine(machine)

			if e != nil {
				return nil, e
			}

			return s.client.VirtualMachineDisks(m)
		},
	)
}
