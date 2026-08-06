package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) RenameDevice(
	name string,
	newName string,
) (*device.Device, error) {
	d, e := c.DeviceByName(name)

	if e != nil {
		return nil, e
	}

	q := netbox.NewPatchedWritableDeviceWithConfigContextRequest()
	q.SetName(newName)

	return c.updateDevice(d, q)
}
