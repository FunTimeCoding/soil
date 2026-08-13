package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetDevicePlatform(
	name string,
	platformName string,
) (*device.Device, error) {
	d, e := c.DeviceByName(name)

	if e != nil {
		return nil, e
	}

	p, f := c.PlatformByName(platformName)

	if f != nil {
		return nil, f
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetPlatform(netbox.DeviceTypeRequestDefaultPlatform{Int32: &p.Identifier})

	return c.updateDevice(d, q)
}
