package service

import "github.com/funtimecoding/soil/pkg/proxmox/network_device"

func (s *Service) DeriveHardwareAddress(
	instance string,
	identifier int,
) (string, *int, error) {
	address, e := network_device.Derive(s.index(instance), identifier)

	if e != nil {
		return "", nil, e
	}

	c, f := s.client(instance)

	if f != nil {
		return "", nil, f
	}

	holder, g := s.addressHolder(c, address, identifier)

	if g != nil {
		return "", nil, g
	}

	return address, holder, nil
}
