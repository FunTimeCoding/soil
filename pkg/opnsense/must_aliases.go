package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/alias"
)

func (c *Client) MustAliases(phrase string) []*alias.Alias {
	result, e := c.Aliases(phrase)
	errors.PanicOnError(e)

	return result
}
