package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/site"

func (c *Client) SitesByMatch(_ string) ([]*site.Site, error) {
	return nil, c.Failure
}
