package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic/float"
	"time"
)

func (c *Client) MustQueryStructFloat(
	q string,
	fallback float64,
	t time.Time,
) *float.Result {
	result, e := c.QueryStructFloat(q, fallback, t)
	errors.PanicOnError(e)

	return result
}
