package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) MustSeries() v1.TSDBResult {
	result, e := c.Series()
	errors.PanicOnError(e)

	return result
}
