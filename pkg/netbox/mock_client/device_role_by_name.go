package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
)

func (c *Client) DeviceRoleByName(n string) (*device_role.Role, error) {
	for _, r := range c.roles {
		if r.Name == n {
			return r, nil
		}
	}

	return nil, not_found.New(constant.DeviceRoleKind, n)
}
