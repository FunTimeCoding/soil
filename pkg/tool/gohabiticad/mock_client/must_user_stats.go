package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/statistic"

func (c *Client) MustUserStats() *statistic.Statistic {
	return c.stats
}
