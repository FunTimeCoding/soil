package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/cluster_group"
)

func (c *Client) MustClusterGroups() []*cluster_group.Group {
	result, e := c.ClusterGroups()
	errors.PanicOnError(e)

	return result
}
