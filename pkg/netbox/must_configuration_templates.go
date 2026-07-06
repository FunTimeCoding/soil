package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/configuration_template"
)

func (c *Client) MustConfigurationTemplates() []*configuration_template.Template {
	result, e := c.ConfigurationTemplates()
	errors.PanicOnError(e)

	return result
}
