package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) DeviceInterfaceExists(
	d *device.Device,
	name string,
) (bool, error) {
	_, e := c.DeviceInterfaceByName(d, name)

	if e != nil {
		if not_found.Is(e) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}
