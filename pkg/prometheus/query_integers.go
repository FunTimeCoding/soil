package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/parse"
	"github.com/funtimecoding/soil/pkg/strings"
	"time"
)

func (c *Client) QueryIntegers(
	q string,
	t time.Time,
) map[string]int {
	result := make(map[string]int)

	for _, r := range parse.Generic(c.MustQuery(q, t).Value) {
		result[r.Metric] = strings.MustToInteger(r.Value)
	}

	return result
}
