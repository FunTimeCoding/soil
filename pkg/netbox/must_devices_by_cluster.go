package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDevicesByCluster(s string) []*device.Device {
	result, e := c.DevicesByCluster(s)
	errors.PanicOnError(e)

	return result
}
