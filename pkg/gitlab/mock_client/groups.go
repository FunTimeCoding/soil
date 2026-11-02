package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Groups(_ context.Context) ([]*gitlab.Group, error) {
	return nil, c.failure
}
