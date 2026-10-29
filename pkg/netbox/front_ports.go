package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/front_port"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) FrontPorts() ([]*front_port.Port, error) {
	result, _, e := c.client.DcimAPI.DcimFrontPortsList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return front_port.NewSlice(result.Results), nil
}
