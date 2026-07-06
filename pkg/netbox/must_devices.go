package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDevices() []*device.Device {
	result, e := c.Devices()
	errors.PanicOnError(e)

	return result
}
