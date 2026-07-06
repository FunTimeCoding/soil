package alertmanager

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustRuleExists(name string) bool {
	result, e := c.RuleExists(name)
	errors.PanicOnError(e)

	return result
}
