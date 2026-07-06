package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/module_type"
)

func (c *Client) MustModuleTypes() []*module_type.Type {
	result, e := c.ModuleTypes()
	errors.PanicOnError(e)

	return result
}
