package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/parse"
	"time"
)

func (c *Client) QueryVector(q string) float64 {
	return parse.VectorFloatSingle(
		c.MustQuery(
			q,
			time.Now().Add(-time.Hour),
		).Value,
	)
}
