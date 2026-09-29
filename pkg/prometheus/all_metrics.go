package prometheus

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
)

func (c *Client) AllMetrics() ([]string, error) {
	result, e := c.LabelValues(constant.Name, []string{}, library.StartOfTime)

	if e != nil {
		return nil, e
	}

	return result.Values, nil
}
