package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) MustSitesByMatch(m string) []*site.Site {
	result, e := c.SitesByMatch(m)
	errors.PanicOnError(e)

	return result
}
