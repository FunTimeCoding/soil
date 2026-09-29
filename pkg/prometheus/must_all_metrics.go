package prometheus

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustAllMetrics() []string {
	result, e := c.AllMetrics()
	errors.PanicOnError(e)

	return result
}
