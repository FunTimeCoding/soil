package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListTunnelGroups(_ context.Context) ([]*client.TunnelGroup, error) {
	return nil, c.Failure
}
