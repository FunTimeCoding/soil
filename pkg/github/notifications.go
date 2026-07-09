package github

import "github.com/google/go-github/v89/github"

func (c *Client) Notifications() ([]*github.Notification, error) {
	result, _, e := c.client.Activity.ListNotifications(
		c.context,
		&github.NotificationListOptions{},
	)

	return result, e
}
