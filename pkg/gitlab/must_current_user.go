package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustCurrentUser() *gitlab.User {
	result, e := c.CurrentUser()
	errors.PanicOnError(e)

	return result
}
