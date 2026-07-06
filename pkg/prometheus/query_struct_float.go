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
) *float.Result {
	result := parse.Generic(c.MustQuery(q, t).Value)

	if len(result) == 0 {
		r := float.New()
		r.Value = fallback

		return r
	}

	r := float.New()
	r.Time = result[0].Time
	r.Value = strings.ToFloat(result[0].Value, fallback)
	r.Raw = result[0]

	return r
}
