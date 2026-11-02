package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Users(_ context.Context) ([]*gitlab.User, error) {
	return nil, c.failure
}
