package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
)

func (c *Client) MustDeviceTypes() []*device_type.Type {
	result, e := c.DeviceTypes()
	errors.PanicOnError(e)

	return result
}
