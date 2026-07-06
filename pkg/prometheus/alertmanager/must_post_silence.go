package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustPostSilence(
	identifier string,
	alert string,
	comment string,
	start time.Time,
	end time.Time,
) string {
	result, e := c.PostSilence(identifier, alert, comment, start, end)
	errors.PanicOnError(e)

	return result
}
