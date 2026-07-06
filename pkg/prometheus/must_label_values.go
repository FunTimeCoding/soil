package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) MustLabelValues(
	label string,
	matches []string,
	since time.Time,
) *label_result.Result {
	result, e := c.LabelValues(label, matches, since)
	errors.PanicOnError(e)

	return result
}
