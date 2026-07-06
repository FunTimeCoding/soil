package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"time"
)

func (c *Client) Query(
	q string,
	t time.Time,
) (*query_result.Result, error) {
	v, w, e := c.client.Query(c.context, q, t)

	if e != nil {
		return nil, e
	}

	return query_result.New(v, w), nil
}
