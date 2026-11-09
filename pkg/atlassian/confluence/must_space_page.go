package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSpacePage(
	spaceName string,
	name string,
) *page.Page {
	result, e := c.SpacePage(spaceName, name)
	errors.PanicOnError(e)

	return result
}
