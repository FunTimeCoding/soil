package habitica

import "github.com/funtimecoding/soil/pkg/habitica/user"

func (c *Client) user() (*user.User, error) {
	var result *user.User
	e := c.basic.Get("/user", nil, &result)

	return result, e
}
