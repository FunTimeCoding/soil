package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/configuration_template"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) ConfigurationTemplates() ([]*configuration_template.Template, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasConfigTemplatesList(c.context).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return configuration_template.NewSlice(result.Results), nil
}
