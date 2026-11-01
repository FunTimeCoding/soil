package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) GroupByName(_ string) (*gitlab.Group, error) {
	return nil, c.failure
}
