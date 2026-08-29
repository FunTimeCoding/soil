package mock_service

import "github.com/funtimecoding/soil/pkg/proxmox/network_device"

func (s *Service) DeriveHardwareAddress(
	instance string,
	identifier int,
) (string, *int, error) {
	address, e := network_device.Derive(s.inventory.Index(instance), identifier)

	if e != nil {
		return "", nil, e
	}

	return address, nil, nil
}
