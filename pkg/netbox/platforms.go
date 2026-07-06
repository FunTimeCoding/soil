package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
)

func (c *Client) Platforms() ([]*platform.Platform, error) {
	result, _, e := c.client.DcimAPI.DcimPlatformsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return platform.NewSlice(result.Results), nil
}
