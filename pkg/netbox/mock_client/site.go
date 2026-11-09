package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) Site(n string) (*site.Site, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, s := range c.sites {
		if s.Name == n {
			return s, nil
		}
	}

	return nil, not_found.New(constant.SiteKind, n)
}
