package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetDeviceDescription(
	name string,
	description string,
) (*device.Device, error) {
	d, e := c.DeviceByName(name)

	if e != nil {
		return nil, e
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetDescription(description)

	return c.updateDevice(d, q)
}
