package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetDeviceSerial(
	name string,
	serial string,
) (*device.Device, error) {
	d, e := c.DeviceByName(name)

	if e != nil {
		return nil, e
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetSerial(serial)

	return c.updateDevice(d, q)
}
