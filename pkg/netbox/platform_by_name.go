package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
)

func (c *Client) PlatformByName(n string) (*platform.Platform, error) {
	result, e := c.Platforms()

	if e != nil {
		return nil, e
	}

	for _, p := range result {
		if p.Name == n {
			return p, nil
		}
	}

	return nil, not_found.New("platform", n)
}
