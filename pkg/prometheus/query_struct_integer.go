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
) (*integer.Result, error) {
	v, e := c.Query(q, t)

	if e != nil {
		return nil, e
	}

	result := parse.Generic(v.Value)
	r := integer.New()

	if len(result) == 0 {
		r.Value = fallback

		return r, nil
	}

	r.Time = result[0].Time
	r.Value = strings.ToInteger(result[0].Value, fallback)
	r.Raw = result[0]

	return r, nil
}
