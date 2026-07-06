package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"time"
)

func (c *Client) MustQuery(
	q string,
	t time.Time,
) *query_result.Result {
	result, e := c.Query(q, t)
	errors.PanicOnError(e)

	return result
}
