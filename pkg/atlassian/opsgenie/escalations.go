package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/escalation"
)

func (c *Client) Escalations() []escalation.Escalation {
	r, e := c.userClient.Escalation.List(c.context)
	errors.PanicOnError(e)

	return r.Escalations
}
