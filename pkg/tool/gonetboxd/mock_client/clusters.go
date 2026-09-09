package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/cluster"

func (c *Client) Clusters() ([]*cluster.Cluster, error) {
	return nil, nil
}
