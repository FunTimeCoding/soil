package netbox

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) DevicesByName(s string) ([]*device.Device, error) {
	result, _, e := c.client.DcimAPI.DcimDevicesList(
		c.context,
	).Name([]string{s}).Execute()

	if e != nil {
		return nil, e
	}

	return device.NewSlice(result.Results), nil
}
