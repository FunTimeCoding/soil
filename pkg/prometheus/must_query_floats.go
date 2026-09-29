package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustQueryFloats(
	q string,
	t time.Time,
) map[string]float64 {
	result, e := c.QueryFloats(q, t)
	errors.PanicOnError(e)

	return result
}
