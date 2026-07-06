package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
)

func (c *Client) PowerOutlets() ([]*power_outlet.Outlet, error) {
	result, _, e := c.client.DcimAPI.DcimPowerOutletsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return power_outlet.NewSlice(result.Results), nil
}
