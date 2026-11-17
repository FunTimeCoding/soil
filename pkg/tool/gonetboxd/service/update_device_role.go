package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) UpdateDeviceRole(
	name string,
	newName *string,
) (*device_role.Role, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*device_role.Role, error) {
			if newName == nil {
				return s.client.DeviceRole(name)
			}

			return s.client.RenameDeviceRole(name, *newName)
		},
	)
}
