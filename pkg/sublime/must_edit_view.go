package sublime

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

func (c *Client) MustEditView(
	identifier int,
	old string,
	new string,
	all bool,
) *view.View {
	result, e := c.EditView(identifier, old, new, all)
	errors.PanicOnError(e)

	return result
}
