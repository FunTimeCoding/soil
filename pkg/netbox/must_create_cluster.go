package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) MustCreateCluster(
	name string,
	t *cluster_type.Type,
	s *site.Site,
) *cluster.Cluster {
	result, e := c.CreateCluster(name, t, s)
	errors.PanicOnError(e)

	return result
}
