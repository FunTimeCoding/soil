package netbox

import "github.com/funtimecoding/soil/pkg/netbox/device_bay_template"

func (c *Client) DeviceBayTemplates() ([]*device_bay_template.Template, error) {
	result, _, e := c.client.DcimAPI.DcimDeviceBayTemplatesList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return device_bay_template.NewSlice(result.Results), nil
}
