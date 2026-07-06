package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) MustLabelNames(
	matches []string,
	since time.Time,
) *label_result.Result {
	result, e := c.LabelNames(matches, since)
	errors.PanicOnError(e)

	return result
}
