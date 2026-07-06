package netbox

import "github.com/funtimecoding/soil/pkg/netbox/cluster"

func (c *Client) Clusters() ([]*cluster.Cluster, error) {
	result, _, e := c.client.VirtualizationAPI.VirtualizationClustersList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return cluster.NewSlice(result.Results), nil
}
