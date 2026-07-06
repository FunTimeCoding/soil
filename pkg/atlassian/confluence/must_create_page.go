package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustCreatePage(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	markdown string,
) *page.Page {
	result, e := c.CreatePage(
		spaceIdentifier,
		parentIdentifier,
		title,
		markdown,
	)
	errors.PanicOnError(e)

	return result
}
