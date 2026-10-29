package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_bay_template"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) DeviceBayTemplates() ([]*device_bay_template.Template, error) {
	result, _, e := c.client.DcimAPI.DcimDeviceBayTemplatesList(c.context).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return device_bay_template.NewSlice(result.Results), nil
}
