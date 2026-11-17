package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) DevicePowerOutlets(
	device int32,
) ([]*power_outlet.Outlet, error) {
	result, _, e := c.client.DcimAPI.DcimPowerOutletsList(c.context).DeviceId(
		[]int32{device},
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return power_outlet.NewSlice(result.Results), nil
}
