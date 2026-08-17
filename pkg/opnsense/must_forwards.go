package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/forward"
)

func (c *Client) MustForwards(phrase string) []*forward.Forward {
	result, e := c.Forwards(phrase)
	errors.PanicOnError(e)

	return result
}
