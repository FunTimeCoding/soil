package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/choice_set"
)

func (c *Client) MustChoiceSetByName(n string) *choice_set.Set {
	result, e := c.ChoiceSetByName(n)
	errors.PanicOnError(e)

	return result
}
