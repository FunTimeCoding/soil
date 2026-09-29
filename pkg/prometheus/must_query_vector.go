package prometheus

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustQueryVector(q string) float64 {
	result, e := c.QueryVector(q)
	errors.PanicOnError(e)

	return result
}
