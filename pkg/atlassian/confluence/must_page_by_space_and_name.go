package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustPageBySpaceAndName(
	spaceName string,
	name string,
) *page.Page {
	result, e := c.PageBySpaceAndName(spaceName, name)
	errors.PanicOnError(e)

	return result
}
