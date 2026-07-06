package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_chassis"
)

func (c *Client) VirtualChassis() ([]*virtual_chassis.Chassis, error) {
	result, _, e := c.client.DcimAPI.DcimVirtualChassisList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return virtual_chassis.NewSlice(result.Results), nil
}
