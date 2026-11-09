package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
)

func (c *Client) MustDeviceRole(n string) *device_role.Role {
	result, e := c.DeviceRole(n)
	errors.PanicOnError(e)

	return result
}
