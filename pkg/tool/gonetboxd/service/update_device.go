package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/types/device_change"
)

func (s *Service) UpdateDevice(
	name string,
	c *device_change.Change,
) (*device.Device, error) {
	if c.PrimaryAddress != nil {
		if _, e := s.client.SetDevicePrimaryAddress(
			name,
			*c.PrimaryAddress,
		); e != nil {
			return nil, e
		}
	}

	if c.Location != nil {
		if _, e := s.client.SetDeviceLocation(name, *c.Location); e != nil {
			return nil, e
		}
	}

	if c.Platform != nil {
		if _, e := s.client.SetDevicePlatform(name, *c.Platform); e != nil {
			return nil, e
		}
	}

	if c.Tenant != nil {
		if _, e := s.client.SetDeviceTenant(name, *c.Tenant); e != nil {
			return nil, e
		}
	}

	if c.Status != nil {
		if _, e := s.client.SetDeviceStatus(name, *c.Status); e != nil {
			return nil, e
		}
	}

	if c.Serial != nil {
		if _, e := s.client.SetDeviceSerial(name, *c.Serial); e != nil {
			return nil, e
		}
	}

	if c.Description != nil {
		if _, e := s.client.SetDeviceDescription(
			name,
			*c.Description,
		); e != nil {
			return nil, e
		}
	}

	if c.Name != nil {
		renamed, e := s.client.RenameDevice(name, *c.Name)

		if e != nil {
			return nil, e
		}

		name = renamed.Name
	}

	return s.client.DeviceByName(name)
}
