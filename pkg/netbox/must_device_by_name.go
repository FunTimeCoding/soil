package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDeviceByName(n string) *device.Device {
	result, e := c.DeviceByName(n)
	errors.PanicOnError(e)

	return result
}
