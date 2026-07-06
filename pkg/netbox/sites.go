package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) Sites() ([]*site.Site, error) {
	if len(c.cache.Sites) != 0 {
		return c.cache.Sites, nil
	}

	result, _, e := c.client.DcimAPI.DcimSitesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.Sites = site.NewSlice(result.Results)

	return c.cache.Sites, nil
}
