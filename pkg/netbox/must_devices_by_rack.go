package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDevicesByRack(i int32) []*device.Device {
	result, e := c.DevicesByRack(i)
	errors.PanicOnError(e)

	return result
}
