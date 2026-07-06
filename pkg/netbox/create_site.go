package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateSite(name string) (*site.Site, error) {
	q := netbox.NewWritableSiteRequest(
		name,
		slug(name),
	)
	result, _, e := c.client.DcimAPI.DcimSitesCreate(
		c.context,
	).WritableSiteRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.Sites = nil

	return site.New(result), nil
}
