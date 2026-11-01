package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) Groups() ([]*gitlab.Group, error) {
	return nil, c.failure
}
