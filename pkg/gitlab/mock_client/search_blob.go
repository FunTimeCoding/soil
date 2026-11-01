package mock_client

import "gitlab.com/gitlab-org/api/client-go/v3"

func (c *Client) SearchBlob(_ string) ([]*gitlab.Blob, error) {
	return nil, c.failure
}
