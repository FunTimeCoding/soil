package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/service_template"
)

func (c *Client) ServiceTemplates() ([]*service_template.Template, error) {
	result, _, e := c.client.IpamAPI.IpamServiceTemplatesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return service_template.NewSlice(result.Results), nil
}
