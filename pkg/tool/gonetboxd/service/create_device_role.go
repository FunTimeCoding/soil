package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateDeviceRole(
	name string,
	virtualMachine bool,
) (*device_role.Role, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*device_role.Role, error) {
			return s.client.CreateDeviceRole(name, virtualMachine)
		},
	)
}
