package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/statistic"

func (c *Client) UserStats() (*statistic.Statistic, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.stats, nil
}
