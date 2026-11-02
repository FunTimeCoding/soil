package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Setting(_ context.Context) (*gitlab.Settings, error) {
	return nil, c.failure
}
