package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedRawDevice(v *netbox.DeviceWithConfigContext) *device.Device {
	v.SetId(c.nextIdentifier())
	result := device.New(v)
	c.devices = append(c.devices, result)

	return result
}
