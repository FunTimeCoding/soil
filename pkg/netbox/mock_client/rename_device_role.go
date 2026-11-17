package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device_role"
)

func (c *Client) RenameDeviceRole(
	name string,
	newName string,
) (*device_role.Role, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for i, r := range c.roles {
		if r.Name != name {
			continue
		}

		c.record("RenameDeviceRole", name, newName)
		r.Raw.SetName(newName)
		c.roles[i] = device_role.New(r.Raw)

		return c.roles[i], nil
	}

	return nil, not_found.New(constant.DeviceRoleKind, name)
}
