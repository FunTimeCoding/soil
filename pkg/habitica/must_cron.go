package habitica

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/habitica/cron"
)

func (c *Client) MustCron() *cron.Cron {
	result, e := c.Cron()
	errors.PanicOnError(e)

	return result
}
