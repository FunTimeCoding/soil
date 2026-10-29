package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/contact_group"

func (c *Client) ContactGroups() ([]*contact_group.Group, error) {
	return nil, c.Failure
}
