package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustPages() []*page.Page {
	result, e := c.Pages()
	errors.PanicOnError(e)

	return result
}
