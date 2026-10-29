package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateClusterType(_ context.Context, _ string) (client.ClusterType, error) {
	return client.ClusterType{}, c.Failure
}
