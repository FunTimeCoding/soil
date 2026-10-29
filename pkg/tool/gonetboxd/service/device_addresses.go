package service

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (s *Service) DeviceAddresses(device string) ([]*internet_address.Address, error) {
	return s.client.DeviceAddresses(device)
}
