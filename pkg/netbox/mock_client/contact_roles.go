package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/contact_role"

func (c *Client) ContactRoles() ([]*contact_role.Role, error) {
	return nil, c.Failure
}
