package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/rule"
)

func (c *Client) MustRule(name string) *rule.Rule {
	result, e := c.Rule(name)
	errors.PanicOnError(e)

	return result
}
