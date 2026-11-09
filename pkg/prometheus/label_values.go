package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/helper"
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) LabelValues(
	label string,
	matches []string,
	start time.Time,
	end time.Time,
) (*label_result.Result, error) {
	v, w, n, e := c.client.LabelValues(c.context, label, matches, start, end)

	if e != nil {
		return nil, e
	}

	return label_result.New(helper.LabelValuesToStrings(v), w, n), nil
}
