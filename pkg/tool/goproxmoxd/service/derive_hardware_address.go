package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/proxmox/network_device"
)

func (s *Service) DeriveHardwareAddress(
	instance string,
	identifier int,
) (string, *int, error) {
	address, e := network_device.Derive(s.index(instance), identifier)

	if e != nil {
		return "", nil, e
	}

	holder, f := guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (*int, error) {
			return s.addressHolder(c, address, identifier)
		},
	)

	if f != nil {
		return "", nil, f
	}

	return address, holder, nil
}
