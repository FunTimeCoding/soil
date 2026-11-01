package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) PackageFiles(_ int64, _ int64) ([]*gitlab.PackageFile, error) {
	return nil, c.failure
}
