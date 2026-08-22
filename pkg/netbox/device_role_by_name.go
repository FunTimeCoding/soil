package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
)

func (c *Client) DeviceRoleByName(n string) (*device_role.Role, error) {
	result, e := c.DeviceRoles()

	if e != nil {
		return nil, e
	}

	for _, r := range result {
		if r.Name == n {
			return r, nil
		}
	}

	return nil, not_found.New("device role", n)
}
