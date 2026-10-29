package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) DeviceInterfaces(device int32) ([]*network.Interface, error) {
	result, _, e := c.client.DcimAPI.DcimInterfacesList(c.context).DeviceId(
		[]int32{device},
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return network.NewSlice(result.Results), nil
}
