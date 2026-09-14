package gitlab

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) Users() ([]*gitlab.User, error) {
	result, _, e := c.client.Users.ListUsers(&gitlab.ListUsersOptions{})

	return result, wrapError(e)
}
