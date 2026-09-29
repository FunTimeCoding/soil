package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustQueryIntegers(
	q string,
	t time.Time,
) map[string]int {
	result, e := c.QueryIntegers(q, t)
	errors.PanicOnError(e)

	return result
}
