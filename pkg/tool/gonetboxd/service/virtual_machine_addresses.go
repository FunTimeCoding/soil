package service

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (s *Service) VirtualMachineAddresses(machine string) ([]*internet_address.Address, error) {
	return s.client.VirtualMachineAddresses(machine)
}
