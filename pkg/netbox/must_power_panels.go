package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/power_panel"
)

func (c *Client) MustPowerPanels() []*power_panel.Panel {
	result, e := c.PowerPanels()
	errors.PanicOnError(e)

	return result
}
