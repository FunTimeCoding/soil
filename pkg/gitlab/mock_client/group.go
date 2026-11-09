package mock_client

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Group(
	_ context.Context,
	_ string,
) (*gitlab.Group, error) {
	return nil, c.failure
}
