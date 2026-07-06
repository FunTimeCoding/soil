package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) MustSiteByName(n string) *site.Site {
	result, e := c.SiteByName(n)
	errors.PanicOnError(e)

	return result
}
