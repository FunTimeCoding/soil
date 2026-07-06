package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
)

func (c *Client) MustModuleBays() []*module_bay.Bay {
	result, e := c.ModuleBays()
	errors.PanicOnError(e)

	return result
}
