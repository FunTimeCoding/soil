package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustCurrentUser(x context.Context) *gitlab.User {
	result, e := c.CurrentUser(x)
	errors.PanicOnError(e)

	return result
}
