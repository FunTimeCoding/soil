package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/service_template"
)

func (c *Client) MustServiceTemplates() []*service_template.Template {
	result, e := c.ServiceTemplates()
	errors.PanicOnError(e)

	return result
}
