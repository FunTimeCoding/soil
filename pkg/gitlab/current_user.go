package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) CurrentUser(x context.Context) (*gitlab.User, error) {
	result, _, e := c.client.Users.CurrentUser(gitlab.WithContext(x))

	return result, wrapError(e)
}
