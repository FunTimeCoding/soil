package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/forwarding_rule"
)

func (c *Client) Forwards() []forwarding_rule.ForwardingRule {
	r, e := c.userClient.Forward.List(
		c.context,
		&forwarding_rule.ListRequest{},
	)
	errors.PanicOnError(e)

	return r.ForwardingRule
}
