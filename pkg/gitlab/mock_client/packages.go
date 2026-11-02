package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Packages(
	_ context.Context,
	_ int64,
	_ bool,
) ([]*gitlab.Package, error) {
	return nil, c.failure
}
