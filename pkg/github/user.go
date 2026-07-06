package github

import "github.com/funtimecoding/soil/pkg/github/user"

func (c *Client) User() (*user.User, error) {
	result, _, e := c.client.Users.Get(c.context, "")

	if e != nil {
		return nil, e
	}

	return user.New(result), nil
}
