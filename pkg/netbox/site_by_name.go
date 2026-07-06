package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (c *Client) SiteByName(n string) (*site.Site, error) {
	result, e := c.SitesByName(n)

	if e != nil {
		return nil, e
	}

	if len(result) != 1 {
		return nil, fmt.Errorf(
			"expected 1 site named %s, got %d",
			n,
			len(result),
		)
	}

	return result[0], nil
}
