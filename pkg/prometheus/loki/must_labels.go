package loki

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustLabels(
	start time.Time,
	end time.Time,
) []string {
	result, e := c.Labels(start, end)
	errors.PanicOnError(e)

	return result
}
