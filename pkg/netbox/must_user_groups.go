package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/user_group"
)

func (c *Client) MustUserGroups() []*user_group.Group {
	result, e := c.UserGroups()
	errors.PanicOnError(e)

	return result
}
