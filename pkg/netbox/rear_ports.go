package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/rear_port"
)

func (c *Client) RearPorts() ([]*rear_port.Port, error) {
	result, _, e := c.client.DcimAPI.DcimRearPortsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return rear_port.NewSlice(result.Results), nil
}
