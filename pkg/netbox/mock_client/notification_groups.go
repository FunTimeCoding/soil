package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/notification_group"

func (c *Client) NotificationGroups() ([]*notification_group.Group, error) {
	return nil, c.Failure
}
