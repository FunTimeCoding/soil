package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListClusters(_ context.Context) ([]*client.Cluster, error) {
	return nil, c.Failure
}
