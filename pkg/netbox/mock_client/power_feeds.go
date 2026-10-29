package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/power_feed"

func (c *Client) PowerFeeds() ([]*power_feed.Feed, error) {
	return nil, c.Failure
}
