package prometheus

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"time"
)

func (c *Client) AllMetrics() ([]string, error) {
	result, e := c.LabelValues(
		constant.Name,
		[]string{},
		soil.Epoch,
		time.Time{},
	)

	if e != nil {
		return nil, e
	}

	return result.Values, nil
}
