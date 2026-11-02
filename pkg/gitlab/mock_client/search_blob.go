package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) SearchBlob(
	_ context.Context,
	_ string,
) ([]*gitlab.Blob, error) {
	return nil, c.failure
}
