package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/rule"
)

func (c *Client) MustRules(phrase string) []*rule.Rule {
	result, e := c.Rules(phrase)
	errors.PanicOnError(e)

	return result
}
