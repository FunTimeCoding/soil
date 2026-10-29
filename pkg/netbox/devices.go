package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) Devices() ([]*device.Device, error) {
	result, _, e := c.client.DcimAPI.DcimDevicesList(c.context).Limit(10000).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return device.NewSlice(result.Results), nil
}
