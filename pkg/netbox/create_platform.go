package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreatePlatform(name string) (*platform.Platform, error) {
	q := netbox.NewPlatformRequest(name, slug(name))
	result, _, e := c.client.DcimAPI.DcimPlatformsCreate(c.context).PlatformRequest(
		*q,
	).Execute()

	if e != nil {
		return nil, e
	}

	return platform.New(result), nil
}
