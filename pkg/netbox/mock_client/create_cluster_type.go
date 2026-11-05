package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/cluster_type"

func (c *Client) CreateClusterType(_ string) (*cluster_type.Type, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
