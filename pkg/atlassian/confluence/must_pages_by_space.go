package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustPagesBySpace(identifier string) []*page.Page {
	result, e := c.PagesBySpace(identifier, "")
	errors.PanicOnError(e)

	return result
}
