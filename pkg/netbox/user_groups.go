package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/user_group"
)

func (c *Client) UserGroups() ([]*user_group.Group, error) {
	result, _, e := c.client.UsersAPI.UsersGroupsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return user_group.NewSlice(result.Results), nil
}
