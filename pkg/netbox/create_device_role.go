package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateDeviceRole(name string) (*device_role.Role, error) {
	q := netbox.NewWritableDeviceRoleRequest(
		name,
		slug(name),
	)
	result, _, e := c.client.DcimAPI.DcimDeviceRolesCreate(
		c.context,
	).WritableDeviceRoleRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.DeviceRoles = nil

	return device_role.New(result), nil
}
