package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SetDeviceRole(
	name string,
	roleName string,
) (*device.Device, error) {
	d, e := c.Device(name)

	if e != nil {
		return nil, e
	}

	r, f := c.DeviceRole(roleName)

	if f != nil {
		return nil, f
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetRole(netbox.DeviceWithConfigContextRequestRole{Int32: &r.Identifier})

	return c.updateDevice(d, q)
}
