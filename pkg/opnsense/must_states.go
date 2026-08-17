package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/state"
)

func (c *Client) MustStates(phrase string) []*state.State {
	result, e := c.States(phrase)
	errors.PanicOnError(e)

	return result
}
