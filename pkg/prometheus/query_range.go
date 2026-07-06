package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) QueryRange(
	q string,
	r v1.Range,
) (*query_result.Result, error) {
	v, w, e := c.client.QueryRange(c.context, q, r)

	if e != nil {
		return nil, e
	}

	return query_result.New(v, w), nil
}
