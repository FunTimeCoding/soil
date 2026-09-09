package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) CreateCluster(
	_ string,
	_ *cluster_type.Type,
	_ *site.Site,
) (*cluster.Cluster, error) {
	return nil, nil
}
