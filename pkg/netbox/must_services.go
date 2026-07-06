package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/service"
)

func (c *Client) MustServices() []*service.Service {
	result, e := c.Services()
	errors.PanicOnError(e)

	return result
}
