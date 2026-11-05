package github

import (
	"github.com/funtimecoding/soil/pkg/github/refusal"
	"github.com/google/go-github/v93/github"
)

func (c *Client) Notifications() ([]*github.Notification, error) {
	result, _, e := c.client.Activity.ListNotifications(
		c.context,
		&github.NotificationListOptions{},
	)

	return result, refusal.Classify(e)
}
