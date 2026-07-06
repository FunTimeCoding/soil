package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDevicesByMatch(s string) []*device.Device {
	result, e := c.DevicesByMatch(s)
	errors.PanicOnError(e)

	return result
}
