package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) Compare(_ int64, _ string, _ string) (*gitlab.Compare, error) {
	return nil, c.failure
}
