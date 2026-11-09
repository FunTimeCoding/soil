package confluence

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (c *Client) spaceDraftPages(spaceIdentifier string) ([]*page.Page, error) {
	return c.DraftPages()
}
