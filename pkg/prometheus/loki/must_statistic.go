package loki

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustStatistic(query string) string {
	result, e := c.Statistic(query)
	errors.PanicOnError(e)

	return result
}
