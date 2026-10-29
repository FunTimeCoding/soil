package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/notification_group"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) NotificationGroups() ([]*notification_group.Group, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasNotificationGroupsList(c.context).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return notification_group.NewSlice(result.Results), nil
}
