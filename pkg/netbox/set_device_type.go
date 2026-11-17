package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetDeviceType(
	name string,
	model string,
) (*device.Device, error) {
	d, e := c.Device(name)

	if e != nil {
		return nil, e
	}

	t, f := c.DeviceType(model)

	if f != nil {
		return nil, f
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetDeviceType(
		netbox.DeviceBayTemplateRequestDeviceType{Int32: &t.Identifier},
	)

	return c.updateDevice(d, q)
}
