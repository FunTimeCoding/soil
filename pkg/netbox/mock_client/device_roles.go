package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device_role"

func (c *Client) DeviceRoles() ([]*device_role.Role, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.roles, nil
}
