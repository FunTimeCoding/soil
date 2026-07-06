package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (c *Client) DraftPage(identifier string) (*page.Page, error) {
	return c.pageWithStatus(identifier, constant.DraftStatus)
}
