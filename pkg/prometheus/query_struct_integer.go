package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/parse"
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic/integer"
	"github.com/funtimecoding/soil/pkg/strings"
	"time"
)

func (c *Client) QueryStructInteger(
	q string,
	fallback int,
	t time.Time,
) *integer.Result {
	result := parse.Generic(c.MustQuery(q, t).Value)

	if len(result) == 0 {
		r := integer.New()
		r.Value = fallback

		return r
	}

	r := integer.New()
	r.Time = result[0].Time
	r.Value = strings.ToInteger(result[0].Value, fallback)
	r.Raw = result[0]

	return r
}
