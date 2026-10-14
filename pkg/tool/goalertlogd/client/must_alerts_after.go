package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/alert"
)

func (c *Client) MustAlertsAfter(
	identifier uint,
	limit int,
) []*alert.Alert {
	result, e := c.AlertsAfter(identifier, limit)
	errors.PanicOnError(e)

	return result
}
