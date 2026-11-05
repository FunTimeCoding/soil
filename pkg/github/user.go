package github

import (
	"github.com/funtimecoding/soil/pkg/github/refusal"
	"github.com/funtimecoding/soil/pkg/github/user"
)

func (c *Client) User() (*user.User, error) {
	result, _, e := c.client.Users.Get(c.context, "")

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return user.New(result), nil
}
