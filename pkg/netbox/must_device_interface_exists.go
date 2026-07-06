package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDeviceInterfaceExists(
	d *device.Device,
	name string,
) bool {
	result, e := c.DeviceInterfaceExists(d, name)
	errors.PanicOnError(e)

	return result
}
