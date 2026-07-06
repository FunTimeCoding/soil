package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/export_template"
)

func (c *Client) MustExportTemplates() []*export_template.Template {
	result, e := c.ExportTemplates()
	errors.PanicOnError(e)

	return result
}
