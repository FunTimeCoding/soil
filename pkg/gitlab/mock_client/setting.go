package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) Setting() (*gitlab.Settings, error) {
	return nil, c.failure
}
