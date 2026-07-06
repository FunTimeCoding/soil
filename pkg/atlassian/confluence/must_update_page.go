package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustUpdatePage(
	identifier string,
	title string,
	markdown string,
	message string,
) *page.Page {
	result, e := c.UpdatePage(identifier, title, markdown, message)
	errors.PanicOnError(e)

	return result
}
