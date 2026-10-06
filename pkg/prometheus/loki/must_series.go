package loki

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustSeries(series string) string {
	result, e := c.Series(series)
	errors.PanicOnError(e)

	return result
}
