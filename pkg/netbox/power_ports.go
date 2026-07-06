package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
)

func (c *Client) PowerPorts() ([]*power_port.Port, error) {
	result, _, e := c.client.DcimAPI.DcimPowerPortsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return power_port.NewSlice(result.Results), nil
}
