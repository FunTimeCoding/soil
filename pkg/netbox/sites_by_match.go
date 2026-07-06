package netbox

import "github.com/funtimecoding/soil/pkg/netbox/site"

func (c *Client) SitesByMatch(m string) ([]*site.Site, error) {
	result, _, e := c.client.DcimAPI.DcimSitesList(
		c.context,
	).NameIc([]string{m}).Execute()

	if e != nil {
		return nil, e
	}

	return site.NewSlice(result.Results), nil
}
