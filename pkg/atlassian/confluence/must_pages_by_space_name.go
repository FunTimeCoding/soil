package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustPagesBySpaceName(n string) []*page.Page {
	result, e := c.PagesBySpaceName(n)
	errors.PanicOnError(e)

	return result
}
