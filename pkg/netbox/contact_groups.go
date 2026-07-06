package netbox

import "github.com/funtimecoding/soil/pkg/netbox/contact_group"

func (c *Client) ContactGroups() ([]*contact_group.Group, error) {
	result, _, e := c.client.TenancyAPI.TenancyContactGroupsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return contact_group.NewSlice(result.Results), nil
}
