package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) MustLabelNames(
	matches []string,
	start time.Time,
	end time.Time,
) *label_result.Result {
	result, e := c.LabelNames(matches, start, end)
	errors.PanicOnError(e)

	return result
}
