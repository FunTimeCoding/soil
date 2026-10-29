package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) SitesByName(n string) ([]*site.Site, error) {
	result, _, e := c.client.DcimAPI.DcimSitesList(c.context).Name([]string{n}).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return site.NewSlice(result.Results), nil
}
