package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) Users() ([]*gitlab.User, error) {
	return nil, c.failure
}
