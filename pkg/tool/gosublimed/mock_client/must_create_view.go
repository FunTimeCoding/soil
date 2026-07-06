package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

func (c *Client) MustCreateView(
	title string,
	content string,
	syntax string,
) *view.View {
	result, e := c.CreateView(title, content, syntax)
	errors.PanicOnError(e)

	return result
}
