package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustAddTag(
	deviceName string,
	tag string,
) *device.Device {
	result, e := c.AddTag(deviceName, tag)
	errors.PanicOnError(e)

	return result
}
