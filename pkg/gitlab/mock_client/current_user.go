package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) CurrentUser() (*gitlab.User, error) {
	return nil, c.failure
}
