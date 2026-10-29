package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListPlatforms(_ context.Context) ([]*client.Platform, error) {
	return nil, c.Failure
}
