package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/console_port"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) ConsolePorts() ([]*console_port.Port, error) {
	result, _, e := c.client.DcimAPI.DcimConsolePortsList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return console_port.NewSlice(result.Results), nil
}
