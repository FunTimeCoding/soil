package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
)

func (c *Client) MustClusterByName(n string) *cluster.Cluster {
	result, e := c.ClusterByName(n)
	errors.PanicOnError(e)

	return result
}
