package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
)

func (c *Client) MustClustersByName(s string) []*cluster.Cluster {
	result, e := c.ClustersByName(s)
	errors.PanicOnError(e)

	return result
}
