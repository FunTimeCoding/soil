package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (s *Service) CreateInterface(
	device string,
	name string,
	interfaceType string,
	address string,
) (*network.Interface, error) {
	var h net.HardwareAddr

	if address != "" {
		var e error
		h, e = physicalAddress(address)

		if e != nil {
			return nil, e
		}
	}

	d, e := s.client.DeviceByName(device)

	if e != nil {
		return nil, e
	}

	result, e := s.client.CreateInterface(
		d,
		name,
		netbox.InterfaceTypeValue(interfaceType),
	)

	if e != nil {
		return nil, e
	}

	if h == nil {
		return result, nil
	}

	if _, f := s.client.CreatePhysicalInterface(h, "", result); f != nil {
		return nil, f
	}

	return result, nil
}
