package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (c *Client) UpdatePageAt(
	identifier string,
	title string,
	markdown string,
	version int,
	message string,
) (*page.Page, error) {
	return c.PutPage(
		identifier,
		title,
		page.ToStorage(markdown),
		version,
		message,
		constant.CurrentStatus,
	)
}
