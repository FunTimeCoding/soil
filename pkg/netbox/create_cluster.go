package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateCluster(
	name string,
	t *cluster_type.Type,
	s *site.Site,
) (*cluster.Cluster, error) {
	q := netbox.NewWritableClusterRequest(
		name,
		netbox.BriefClusterTypeRequestAsClusterRequestType(
			netbox.NewBriefClusterTypeRequest(t.Name, t.Raw.Slug),
		),
	)
	q.SetScopeType("dcim.site")
	q.SetScopeId(s.Identifier)
	q.SetStatus(netbox.CLUSTERSTATUSVALUE_ACTIVE)
	result, _, e := c.client.VirtualizationAPI.VirtualizationClustersCreate(
		c.context,
	).WritableClusterRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return cluster.New(result), nil
}
