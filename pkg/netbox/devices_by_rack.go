package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) DevicesByRack(i int32) ([]*device.Device, error) {
	result, _, e := c.client.DcimAPI.DcimDevicesList(c.context).RackId(
		[]int32{i},
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return device.NewSlice(result.Results), nil
}
