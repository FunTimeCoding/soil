package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListTunnelTerminations(_ context.Context) ([]*client.TunnelTermination, error) {
	return nil, c.Failure
}
