package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic/integer"
	"time"
)

func (c *Client) MustQueryStructInteger(
	q string,
	fallback int,
	t time.Time,
) *integer.Result {
	result, e := c.QueryStructInteger(q, fallback, t)
	errors.PanicOnError(e)

	return result
}
