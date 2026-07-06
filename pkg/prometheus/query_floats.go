package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/parse"
	"github.com/funtimecoding/soil/pkg/strings"
	"time"
)

func (c *Client) QueryFloats(
	q string,
	t time.Time,
) map[string]float64 {
	result := make(map[string]float64)

	for _, r := range parse.Generic(c.MustQuery(q, t).Value) {
		result[r.Metric] = strings.MustToFloat(r.Value)
	}

	return result
}
