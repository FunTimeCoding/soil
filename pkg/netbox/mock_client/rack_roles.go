package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/rack_role"

func (c *Client) RackRoles() ([]*rack_role.Role, error) {
	return nil, c.Failure
}
