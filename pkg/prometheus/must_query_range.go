package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) MustQueryRange(
	q string,
	r v1.Range,
) *query_result.Result {
	result, e := c.QueryRange(q, r)
	errors.PanicOnError(e)

	return result
}
