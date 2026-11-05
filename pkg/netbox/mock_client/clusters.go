package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/cluster"

func (c *Client) Clusters() ([]*cluster.Cluster, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
