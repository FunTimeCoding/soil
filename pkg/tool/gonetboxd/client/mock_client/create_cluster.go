package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateCluster(_ context.Context, _ string, _ string, _ string) (client.Cluster, error) {
	return client.Cluster{}, c.Failure
}
