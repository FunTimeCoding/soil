package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) DevicePowerPorts(device int32) ([]*power_port.Port, error) {
	result, _, e := c.client.DcimAPI.DcimPowerPortsList(c.context).DeviceId(
		[]int32{device},
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return power_port.NewSlice(result.Results), nil
}
