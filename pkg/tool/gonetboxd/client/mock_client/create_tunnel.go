package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateTunnel(_ context.Context, _ string, _ string, _ string) (client.Tunnel, error) {
	return client.Tunnel{}, c.Failure
}
