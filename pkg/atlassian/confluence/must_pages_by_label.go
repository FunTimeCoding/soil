package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustPagesByLabel(labelIdentifier string) []*page.Page {
	result, e := c.PagesByLabel(labelIdentifier)
	errors.PanicOnError(e)

	return result
}
