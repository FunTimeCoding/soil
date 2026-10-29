package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/power_feed"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) PowerFeeds() ([]*power_feed.Feed, error) {
	result, _, e := c.client.DcimAPI.DcimPowerFeedsList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return power_feed.NewSlice(result.Results), nil
}
