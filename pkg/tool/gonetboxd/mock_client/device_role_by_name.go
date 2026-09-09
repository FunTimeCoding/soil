package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device_role"

func (c *Client) DeviceRoleByName(_ string) (*device_role.Role, error) {
	return nil, nil
}
