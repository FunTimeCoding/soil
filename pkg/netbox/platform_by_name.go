package netbox

import (
	"fmt"
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

	return nil, fmt.Errorf("platform %s not found", n)
}
