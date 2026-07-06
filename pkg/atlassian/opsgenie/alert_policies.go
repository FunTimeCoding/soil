package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/policy"
)

func (c *Client) AlertPolicies(team string) []policy.PolicyProps {
	r, e := c.userClient.Policy.ListAlertPolicies(
		c.context,
		&policy.ListAlertPoliciesRequest{TeamId: team},
	)
	errors.PanicOnError(e)

	return r.Policies
}
