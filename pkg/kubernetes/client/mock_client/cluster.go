package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/constant"

func (c *Client) Cluster() string {
	return constant.InCluster
}
