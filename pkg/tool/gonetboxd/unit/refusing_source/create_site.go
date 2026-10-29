package refusing_source

import "github.com/funtimecoding/soil/pkg/netbox/site"

func (c *Client) CreateSite(_ string) (*site.Site, error) {
	return nil, c.refusal
}
