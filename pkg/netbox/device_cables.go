package netbox

import "github.com/funtimecoding/soil/pkg/netbox/cable"

func (c *Client) DeviceCables(device string) ([]*cable.Cable, error) {
	result, _, e := c.client.DcimAPI.DcimCablesList(
		c.context,
	).Device([]string{device}).Execute()

	if e != nil {
		return nil, e
	}

	return cable.NewSlice(result.Results), nil
}
