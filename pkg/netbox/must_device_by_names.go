package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDeviceByNames(n []string) *device.Device {
	result, e := c.DeviceByNames(n)
	errors.PanicOnError(e)

	return result
}
