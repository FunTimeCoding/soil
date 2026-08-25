package prometheus

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func (c *Client) AllMetrics() []string {
	return c.MustLabelValues(
		constant.Name,
		[]string{},
		library.StartOfTime,
	).Values
}
