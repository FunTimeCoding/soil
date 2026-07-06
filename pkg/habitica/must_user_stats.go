package habitica

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
)

func (c *Client) MustUserStats() *statistic.Statistic {
	result, e := c.UserStats()
	errors.PanicOnError(e)

	return result
}
