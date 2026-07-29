package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
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
		constant.ConfluenceCurrentStatus,
	)
}
