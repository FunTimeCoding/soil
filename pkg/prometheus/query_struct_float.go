package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/parse"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic/float"
	"github.com/funtimecoding/soil/pkg/strings"
	"time"
)

func (c *Client) QueryStructFloat(
	q string,
	fallback float64,
	t time.Time,
) (*float.Result, error) {
	v, e := c.Query(q, t)

	if e != nil {
		return nil, e
	}

	result := parse.Generic(v.Value)
	r := float.New()

	if len(result) == 0 {
		r.Value = fallback

		return r, nil
	}

	r.Time = result[0].Time
	r.Value = strings.ToFloat(result[0].Value, fallback)
	r.Raw = result[0]

	return r, nil
}
