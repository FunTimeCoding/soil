package sublime

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

func (c *Client) MustOpenFile(path string) *view.View {
	result, e := c.OpenFile(path)
	errors.PanicOnError(e)

	return result
}
