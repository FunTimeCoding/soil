package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_type"
)

func (c *Client) MustClusterTypeByName(n string) *cluster_type.Type {
	result, e := c.ClusterTypeByName(n)
	errors.PanicOnError(e)

	return result
}
