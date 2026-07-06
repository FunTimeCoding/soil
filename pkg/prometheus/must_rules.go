package prometheus

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
)

func (c *Client) MustRules() *rule_list.List {
	result, e := c.Rules()
	errors.PanicOnError(e)

	return result
}
