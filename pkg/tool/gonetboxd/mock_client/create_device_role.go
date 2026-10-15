package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/device_role"

func (c *Client) CreateDeviceRole(
	_ string,
	_ bool,
) (*device_role.Role, error) {
	return nil, nil
}
