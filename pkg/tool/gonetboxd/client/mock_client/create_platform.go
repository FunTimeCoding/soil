package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreatePlatform(_ context.Context, _ string) (client.Platform, error) {
	return client.Platform{}, c.Failure
}
