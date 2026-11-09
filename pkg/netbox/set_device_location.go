package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetDeviceLocation(
	name string,
	locationName string,
) (*device.Device, error) {
	d, e := c.Device(name)

	if e != nil {
		return nil, e
	}

	l, f := c.Location(locationName)

	if f != nil {
		return nil, f
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetLocation(
		netbox.DeviceWithConfigContextRequestLocation{Int32: &l.Identifier},
	)

	return c.updateDevice(d, q)
}
