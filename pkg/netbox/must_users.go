package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/user"
)

func (c *Client) MustUsers() []*user.User {
	result, e := c.Users()
	errors.PanicOnError(e)

	return result
}
