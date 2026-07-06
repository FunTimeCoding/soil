package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/user"
)

func (c *Client) MustUser() *user.User {
	result, e := c.User()
	errors.PanicOnError(e)

	return result
}
