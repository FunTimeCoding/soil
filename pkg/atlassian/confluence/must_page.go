package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustPage(identifier string) *page.Page {
	result, e := c.Page(identifier)
	errors.PanicOnError(e)

	return result
}
