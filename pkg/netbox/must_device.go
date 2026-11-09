package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustDevice(n string) *device.Device {
	result, e := c.Device(n)
	errors.PanicOnError(e)

	return result
}
