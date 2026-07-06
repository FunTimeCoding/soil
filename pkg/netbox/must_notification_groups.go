package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/notification_group"
)

func (c *Client) MustNotificationGroups() []*notification_group.Group {
	result, e := c.NotificationGroups()
	errors.PanicOnError(e)

	return result
}
