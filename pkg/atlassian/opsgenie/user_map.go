package opsgenie

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user_map"

func (c *Client) UserMap() *user_map.Map {
	if c.userMap == nil {
		c.userMap = user_map.New(c.Users())
	}

	return c.userMap
}
