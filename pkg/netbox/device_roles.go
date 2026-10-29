package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) DeviceRoles() ([]*device_role.Role, error) {
	if len(c.cache.DeviceRoles) != 0 {
		return c.cache.DeviceRoles, nil
	}

	result, _, e := c.client.DcimAPI.DcimDeviceRolesList(c.context).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	c.cache.DeviceRoles = device_role.NewSlice(result.Results)

	return c.cache.DeviceRoles, nil
}
