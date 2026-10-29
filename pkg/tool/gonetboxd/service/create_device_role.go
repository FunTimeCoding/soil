package service

import "github.com/funtimecoding/soil/pkg/netbox/device_role"

func (s *Service) CreateDeviceRole(
	name string,
	virtualMachine bool,
) (*device_role.Role, error) {
	return s.client.CreateDeviceRole(name, virtualMachine)
}
