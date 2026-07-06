package netbox

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) DeleteInterfaceByName(
	d *device.Device,
	name string,
) error {
	i, e := c.DeviceInterfaceByName(d, name)

	if e != nil {
		return e
	}

	return c.DeleteInterface(i.Identifier)
}
