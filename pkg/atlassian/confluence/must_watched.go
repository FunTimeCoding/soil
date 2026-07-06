package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustWatched() []*page.Page {
	result, e := c.Watched()
	errors.PanicOnError(e)

	return result
}
