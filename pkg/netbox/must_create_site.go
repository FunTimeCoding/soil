package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) MustCreateSite(name string) *site.Site {
	result, e := c.CreateSite(name)
	errors.PanicOnError(e)

	return result
}
