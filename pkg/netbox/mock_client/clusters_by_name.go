package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/cluster"

func (c *Client) ClustersByName(_ string) ([]*cluster.Cluster, error) {
	return nil, c.Failure
}
