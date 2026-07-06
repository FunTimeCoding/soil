package confluence

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (c *Client) Page(identifier string) (*page.Page, error) {
	return c.pageWithStatus(identifier, "")
}
