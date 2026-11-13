package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_disk"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateVirtualDisk(
	machine string,
	name string,
	size int32,
) (*virtual_disk.Disk, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*virtual_disk.Disk, error) {
			return s.client.CreateVirtualDisk(machine, name, size)
		},
	)
}
