package netbox

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) DeviceInterfaceExists(
	d *device.Device,
	name string,
) (bool, error) {
	_, e := c.DeviceInterfaceByName(d, name)

	if e != nil {
		if errors.Is(e, constant.ErrorNotFound) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}
