package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) SeedDetail(v *gitlab.MergeRequest) {
	c.details = append(c.details, v)
}
