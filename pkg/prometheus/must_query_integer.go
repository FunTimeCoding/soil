package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustQueryInteger(
	q string,
	t time.Time,
) int {
	result, e := c.QueryInteger(q, t)
	errors.PanicOnError(e)

	return result
}
