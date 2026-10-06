package loki

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/loki/basic/query_result"
)

func (c *Client) MustQuery(query string) *query_result.Result {
	result, e := c.Query(query)
	errors.PanicOnError(e)

	return result
}
