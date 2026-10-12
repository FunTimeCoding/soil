package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) SeedDevice(name string) *device.Device {
	result := newDevice(c.nextIdentifier(), name, nil, nil)
	c.devices = append(c.devices, result)

	return result
}
