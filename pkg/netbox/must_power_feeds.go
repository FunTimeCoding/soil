package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/power_feed"
)

func (c *Client) MustPowerFeeds() []*power_feed.Feed {
	result, e := c.PowerFeeds()
	errors.PanicOnError(e)

	return result
}
