package netbox

import "github.com/funtimecoding/soil/pkg/netbox/notification_group"

func (c *Client) NotificationGroups() ([]*notification_group.Group, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasNotificationGroupsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return notification_group.NewSlice(result.Results), nil
}
