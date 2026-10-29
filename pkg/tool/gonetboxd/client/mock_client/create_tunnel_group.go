package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateTunnelGroup(_ context.Context, _ string) (client.TunnelGroup, error) {
	return client.TunnelGroup{}, c.Failure
}
