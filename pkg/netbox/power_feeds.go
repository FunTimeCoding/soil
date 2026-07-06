package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/power_feed"
)

func (c *Client) PowerFeeds() ([]*power_feed.Feed, error) {
	result, _, e := c.client.DcimAPI.DcimPowerFeedsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return power_feed.NewSlice(result.Results), nil
}
