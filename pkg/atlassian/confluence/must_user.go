package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/user"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustUser() *user.User {
	result, e := c.User()
	errors.PanicOnError(e)

	return result
}
