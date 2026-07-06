package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/contact_role"
)

func (c *Client) MustContactRoles() []*contact_role.Role {
	result, e := c.ContactRoles()
	errors.PanicOnError(e)

	return result
}
