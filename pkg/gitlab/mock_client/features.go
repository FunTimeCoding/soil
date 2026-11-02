package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Features(_ context.Context) ([]*gitlab.Feature, error) {
	return nil, c.failure
}
