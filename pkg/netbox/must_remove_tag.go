package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) MustRemoveTag(
	deviceName string,
	tag string,
) *device.Device {
	result, e := c.RemoveTag(deviceName, tag)
	errors.PanicOnError(e)

	return result
}
