package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedSite(name string) {
	v := netbox.NewSiteWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetName(name)
	c.sites = append(c.sites, site.New(v))
}
