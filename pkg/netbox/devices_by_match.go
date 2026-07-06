package netbox

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (c *Client) DevicesByMatch(s string) ([]*device.Device, error) {
	result, _, e := c.client.DcimAPI.DcimDevicesList(
		c.context,
	).NameIc([]string{s}).Execute()

	if e != nil {
		return nil, e
	}

	return device.NewSlice(result.Results), nil
}
