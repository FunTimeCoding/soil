package sublime

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

func (c *Client) MustView(identifier int) *view.View {
	result, e := c.View(identifier)
	errors.PanicOnError(e)

	return result
}
