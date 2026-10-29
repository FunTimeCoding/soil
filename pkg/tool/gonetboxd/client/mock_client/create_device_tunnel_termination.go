package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateDeviceTunnelTermination(_ context.Context, _ string, _ string, _ string, _ string) (client.TunnelTermination, error) {
	return client.TunnelTermination{}, c.Failure
}
