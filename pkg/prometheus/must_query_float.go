package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustQueryFloat(
	q string,
	t time.Time,
) float64 {
	result, e := c.QueryFloat(q, t)
	errors.PanicOnError(e)

	return result
}
