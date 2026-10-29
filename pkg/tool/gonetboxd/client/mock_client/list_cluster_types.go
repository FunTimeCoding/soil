package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListClusterTypes(_ context.Context) ([]*client.ClusterType, error) {
	return nil, c.Failure
}
