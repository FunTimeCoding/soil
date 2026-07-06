package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/module"
)

func (c *Client) MustModules() []*module.Module {
	result, e := c.Modules()
	errors.PanicOnError(e)

	return result
}
