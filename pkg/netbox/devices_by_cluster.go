package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) DevicesByCluster(s string) ([]*device.Device, error) {
	clusters, e := c.ClustersByName(s)

	if e != nil {
		return nil, e
	}

	if len(clusters) == 0 {
		return nil, not_found.New("cluster", s)
	}

	if len(clusters) > 1 {
		return nil, ambiguous.Format(
			"expected 1 cluster for %s, got %d",
			s,
			len(clusters),
		)
	}

	result, _, f := c.client.DcimAPI.DcimDevicesList(c.context).ClusterId(
		[]*int32{&clusters[0].Identifier},
	).Execute()

	if f != nil {
		return nil, f
	}

	return device.NewSlice(result.Results), nil
}
