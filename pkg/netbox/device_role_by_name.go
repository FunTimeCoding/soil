package netbox

import (
	"fmt"
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

	return nil, fmt.Errorf("device role not found: %s", n)
}
