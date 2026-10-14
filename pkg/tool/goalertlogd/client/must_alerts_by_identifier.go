package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/alert"
)

func (c *Client) MustAlertsByIdentifier(identifiers []uint) []*alert.Alert {
	result, e := c.AlertsByIdentifier(identifiers)
	errors.PanicOnError(e)

	return result
}
