package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) MustSites() []*site.Site {
	result, e := c.Sites()
	errors.PanicOnError(e)

	return result
}
