package loki

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustLabelValues(
	start time.Time,
	end time.Time,
	label string,
) []string {
	result, e := c.LabelValues(start, end, label)
	errors.PanicOnError(e)

	return result
}
