package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/notification"
)

func (c *Client) NotificationRuleSteps(
	user string,
	rule string,
) []notification.RuleStep {
	r, e := c.userClient.Notification.ListRuleStep(
		c.context,
		&notification.ListRuleStepsRequest{
			UserIdentifier: user,
			RuleId:         rule,
		},
	)
	errors.PanicOnError(e)

	return r.RuleSteps
}
