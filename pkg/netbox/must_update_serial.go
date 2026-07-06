package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustUpdateSerial(
	device string,
	serial string,
) *device.Device {
	result, e := c.UpdateSerial(device, serial)
	errors.PanicOnError(e)

	return result
}
