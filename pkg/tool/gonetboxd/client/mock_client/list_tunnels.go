package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListTunnels(_ context.Context) ([]*client.Tunnel, error) {
	return nil, c.Failure
}
