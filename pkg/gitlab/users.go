package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Users(x context.Context) ([]*gitlab.User, error) {
	result, _, e := c.client.Users.ListUsers(
		&gitlab.ListUsersOptions{},
		gitlab.WithContext(x),
	)

	return result, wrapError(e)
}
