package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) PackageFiles(
	_ context.Context,
	_ int64,
	_ int64,
) ([]*gitlab.PackageFile, error) {
	return nil, c.failure
}
