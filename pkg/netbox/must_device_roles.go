package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
)

func (c *Client) MustDeviceRoles() []*device_role.Role {
	result, e := c.DeviceRoles()
	errors.PanicOnError(e)

	return result
}
