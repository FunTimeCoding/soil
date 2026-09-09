package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/cluster"

func (c *Client) ClusterByName(_ string) (*cluster.Cluster, error) {
	return nil, nil
}
