package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) MustSitesByName(n string) []*site.Site {
	result, e := c.SitesByName(n)
	errors.PanicOnError(e)

	return result
}
