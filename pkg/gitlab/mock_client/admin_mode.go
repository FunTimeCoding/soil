package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) AdminMode(_ bool) (*gitlab.Settings, error) {
	return nil, c.failure
}
