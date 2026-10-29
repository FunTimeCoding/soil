package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedAddressedDevice(
	name string,
	address string,
	tags ...string,
) *device.Device {
	v := newDevice(c.nextIdentifier(), name, tags, nil).Raw
	v.SetPrimaryIp4(
		*netbox.NewBriefIPAddress(
			0,
			"",
			address,
			*netbox.NewAggregateFamilyWithDefaults(),
			address,
		),
	)
	result := device.New(v)
	c.devices = append(c.devices, result)

	return result
}
