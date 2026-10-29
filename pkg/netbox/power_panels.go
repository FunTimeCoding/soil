package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/power_panel"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) PowerPanels() ([]*power_panel.Panel, error) {
	result, _, e := c.client.DcimAPI.DcimPowerPanelsList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return power_panel.NewSlice(result.Results), nil
}
