package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) DraftPage(identifier string) (*page.Page, error) {
	return c.pageWithStatus(identifier, constant.ConfluenceDraftStatus)
}
