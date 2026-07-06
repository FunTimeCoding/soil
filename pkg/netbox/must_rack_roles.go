package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/rack_role"
)

func (c *Client) MustRackRoles() []*rack_role.Role {
	result, e := c.RackRoles()
	errors.PanicOnError(e)

	return result
}
