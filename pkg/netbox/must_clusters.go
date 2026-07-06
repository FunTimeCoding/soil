package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cluster"
)

func (c *Client) MustClusters() []*cluster.Cluster {
	result, e := c.Clusters()
	errors.PanicOnError(e)

	return result
}
