package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustLabelPages(labelIdentifier string) []*page.Page {
	result, e := c.LabelPages(labelIdentifier)
	errors.PanicOnError(e)

	return result
}
