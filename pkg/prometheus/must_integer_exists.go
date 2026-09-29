package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustIntegerExists(
	q string,
	t time.Time,
) bool {
	result, e := c.IntegerExists(q, t)
	errors.PanicOnError(e)

	return result
}
