package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
)

func (c *Client) MustDeviceRoleByName(n string) *device_role.Role {
	result, e := c.DeviceRoleByName(n)
	errors.PanicOnError(e)

	return result
}
