package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/user"
)

func (c *Client) Users() ([]*user.User, error) {
	result, _, e := c.client.UsersAPI.UsersUsersList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return user.NewSlice(result.Results), nil
}
