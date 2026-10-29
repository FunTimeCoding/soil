package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/export_template"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) ExportTemplates() ([]*export_template.Template, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasExportTemplatesList(c.context).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return export_template.NewSlice(result.Results), nil
}
