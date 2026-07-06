package netbox

import "github.com/funtimecoding/soil/pkg/netbox/cluster_type"

func (c *Client) ClusterTypes() ([]*cluster_type.Type, error) {
	result, _, e := c.client.VirtualizationAPI.VirtualizationClusterTypesList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return cluster_type.NewSlice(result.Results), nil
}
