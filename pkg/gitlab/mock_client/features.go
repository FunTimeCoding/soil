package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) Features() ([]*gitlab.Feature, error) {
	return nil, c.failure
}
