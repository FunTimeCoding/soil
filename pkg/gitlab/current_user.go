package gitlab

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) CurrentUser() (*gitlab.User, error) {
	result, _, e := c.client.Users.CurrentUser()

	return result, wrapError(e)
}
