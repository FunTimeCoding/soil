package netbox

import "github.com/funtimecoding/soil/pkg/netbox/network"

func (c *Client) DeviceInterfaces(device int32) ([]*network.Interface, error) {
	result, _, e := c.client.DcimAPI.DcimInterfacesList(
		c.context,
	).DeviceId([]int32{device}).Execute()

	if e != nil {
		return nil, e
	}

	return network.NewSlice(result.Results), nil
}
