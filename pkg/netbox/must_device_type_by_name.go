package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
)

func (c *Client) MustDeviceTypeByName(n string) *device_type.Type {
	result, e := c.DeviceTypeByName(n)
	errors.PanicOnError(e)

	return result
}
