package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/user_group"

func (c *Client) UserGroups() ([]*user_group.Group, error) {
	return nil, c.Failure
}
