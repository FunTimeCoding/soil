package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) MustSite(n string) *site.Site {
	result, e := c.Site(n)
	errors.PanicOnError(e)

	return result
}
