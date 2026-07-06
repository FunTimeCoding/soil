package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustUnsetSerial(device string) *device.Device {
	result, e := c.UnsetSerial(device)
	errors.PanicOnError(e)

	return result
}
