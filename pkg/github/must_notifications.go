package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/google/go-github/v89/github"
)

func (c *Client) MustNotifications() []*github.Notification {
	result, e := c.Notifications()
	errors.PanicOnError(e)

	return result
}
