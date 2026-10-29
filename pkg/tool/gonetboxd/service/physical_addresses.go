package service

import "github.com/funtimecoding/soil/pkg/netbox/physical_address"

func (s *Service) PhysicalAddresses() ([]*physical_address.Address, error) {
	return s.client.PhysicalAddresses()
}
