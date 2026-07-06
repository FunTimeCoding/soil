package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/device_bay_template"
)

func (c *Client) MustDeviceBayTemplates() []*device_bay_template.Template {
	result, e := c.DeviceBayTemplates()
	errors.PanicOnError(e)

	return result
}
