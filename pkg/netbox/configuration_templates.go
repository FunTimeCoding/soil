package netbox

import "github.com/funtimecoding/soil/pkg/netbox/configuration_template"

func (c *Client) ConfigurationTemplates() ([]*configuration_template.Template, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasConfigTemplatesList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return configuration_template.NewSlice(result.Results), nil
}
