package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/parse"
	"github.com/funtimecoding/soil/pkg/strings"
	"time"
)

func (c *Client) QueryFloats(
	q string,
	t time.Time,
) (map[string]float64, error) {
	result := make(map[string]float64)
	v, e := c.Query(q, t)

	if e != nil {
		return nil, e
	}

	for _, r := range parse.Generic(v.Value) {
		value, f := strings.ParseFloat(r.Value)

		if f != nil {
			return nil, f
		}

		result[r.Metric] = value
	}

	return result, nil
}
