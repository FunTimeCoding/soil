package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) AdminMode(
	_ context.Context,
	_ bool,
) (*gitlab.Settings, error) {
	return nil, c.failure
}
