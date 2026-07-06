package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDevicesByName(s string) []*device.Device {
	result, e := c.DevicesByName(s)
	errors.PanicOnError(e)

	return result
}
