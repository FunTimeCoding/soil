package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/contact_role"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) ContactRoles() ([]*contact_role.Role, error) {
	result, _, e := c.client.TenancyAPI.TenancyContactRolesList(c.context).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return contact_role.NewSlice(result.Results), nil
}
