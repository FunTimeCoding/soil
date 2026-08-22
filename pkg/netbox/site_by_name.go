package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) SiteByName(n string) (*site.Site, error) {
	result, e := c.SitesByName(n)

	if e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return nil, not_found.New("site", n)
	}

	if len(result) > 1 {
		return nil, ambiguous.Format(
			"expected 1 site named %s, got %d",
			n,
			len(result),
		)
	}

	return result[0], nil
}
