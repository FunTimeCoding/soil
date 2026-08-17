package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/source_nat"
)

func (c *Client) MustSourceNatRules(phrase string) []*source_nat.Rule {
	result, e := c.SourceNatRules(phrase)
	errors.PanicOnError(e)

	return result
}
