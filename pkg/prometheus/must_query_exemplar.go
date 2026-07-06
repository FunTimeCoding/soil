package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

func (c *Client) MustQueryExemplar(
	q string,
	start time.Time,
	end time.Time,
) []v1.ExemplarQueryResult {
	result, e := c.QueryExemplar(q, start, end)
	errors.PanicOnError(e)

	return result
}
