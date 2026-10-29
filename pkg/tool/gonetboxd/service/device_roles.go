package service

import "github.com/funtimecoding/soil/pkg/netbox/device_role"

func (s *Service) DeviceRoles() ([]*device_role.Role, error) {
	return s.client.DeviceRoles()
}
