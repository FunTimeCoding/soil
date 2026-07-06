package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateClusterType(name string) (*cluster_type.Type, error) {
	q := netbox.NewClusterTypeRequest(
		name,
		slug(name),
	)
	result, _, e := c.client.VirtualizationAPI.VirtualizationClusterTypesCreate(
		c.context,
	).ClusterTypeRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return cluster_type.New(result), nil
}
