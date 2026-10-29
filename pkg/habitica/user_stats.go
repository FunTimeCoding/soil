package habitica

import "github.com/funtimecoding/soil/pkg/habitica/statistic"

func (c *Client) UserStats() (*statistic.Statistic, error) {
	result, e := c.user()

	if e != nil {
		return nil, e
	}

	return result.Stats, nil
}
