package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) ClustersByName(s string) ([]*cluster.Cluster, error) {
	result, _, e := c.client.VirtualizationAPI.VirtualizationClustersList(
		c.context,
	).NameIc([]string{s}).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return cluster.NewSlice(result.Results), nil
}
