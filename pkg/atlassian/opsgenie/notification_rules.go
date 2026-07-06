package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/notification"
)

func (c *Client) NotificationRules(user string) []notification.SimpleNotificationRuleResult {
	r, e := c.userClient.Notification.ListRule(
		c.context,
		&notification.ListRuleRequest{UserIdentifier: user},
	)
	errors.PanicOnError(e)

	return r.SimpleNotificationRules
}
