package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) MustSetSilence(
	alert string,
	comment string,
	d time.Duration,
) string {
	result, e := c.SetSilence(alert, comment, d)
	errors.PanicOnError(e)

	return result
}
