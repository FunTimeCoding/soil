package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
)

func (c *Client) MustCreateDeviceRole(
	name string,
	virtualMachine bool,
) *device_role.Role {
	result, e := c.CreateDeviceRole(name, virtualMachine)
	errors.PanicOnError(e)

	return result
}
