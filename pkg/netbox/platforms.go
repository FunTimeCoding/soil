package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/platform"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) Platforms() ([]*platform.Platform, error) {
	result, _, e := c.client.DcimAPI.DcimPlatformsList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return platform.NewSlice(result.Results), nil
}
