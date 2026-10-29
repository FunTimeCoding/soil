package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster_group"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) ClusterGroups() ([]*cluster_group.Group, error) {
	result, _, e := c.client.VirtualizationAPI.VirtualizationClusterGroupsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return cluster_group.NewSlice(result.Results), nil
}
