package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/helper"
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) LabelValues(
	label string,
	matches []string,
	since time.Time,
) (*label_result.Result, error) {
	v, w, e := c.client.LabelValues(
		c.context,
		label,
		matches,
		since,
		time.Now(),
	)

	if e != nil {
		return nil, e
	}

	return label_result.New(helper.LabelValuesToStrings(v), w), nil
}
