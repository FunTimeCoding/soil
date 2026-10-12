package service

import "github.com/netbox-community/go-netbox/v4"

func (s *Service) SetVirtualInterfacePhysicalAddress(
	machine string,
	interfaceName string,
	address string,
) (*netbox.VMInterface, error) {
	h, e := physicalAddress(address)

	if e != nil {
		return nil, e
	}

	m, e := s.client.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	return s.client.UpdateVirtualInterface(m, interfaceName, h)
}
