package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustImport(
	space string,
	parent string,
	base string,
	name string,
) *page.Page {
	result, e := c.Import(space, parent, base, name)
	errors.PanicOnError(e)

	return result
}
