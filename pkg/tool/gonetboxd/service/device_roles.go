package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) DeviceRoles() ([]*device_role.Role, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*device_role.Role, error) {
			return s.client.DeviceRoles()
		},
	)
}
