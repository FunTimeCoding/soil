package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

func (c *Client) MustViews() []*view.View {
	result, e := c.Views()
	errors.PanicOnError(e)

	return result
}
