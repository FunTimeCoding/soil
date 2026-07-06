package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustLabeled() []*page.Page {
	result, e := c.Labeled()
	errors.PanicOnError(e)

	return result
}
