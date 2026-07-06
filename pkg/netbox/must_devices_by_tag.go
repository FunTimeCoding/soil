package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDevicesByTag(s string) []*device.Device {
	result, e := c.DevicesByTag(s)
	errors.PanicOnError(e)

	return result
}
