package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustUsers() []*gitlab.User {
	result, e := c.Users()
	errors.PanicOnError(e)

	return result
}
