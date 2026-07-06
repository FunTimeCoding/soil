package confluence

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (c *Client) PagesBySpaceName(n string) ([]*page.Page, error) {
	s, e := c.SpaceByName(n)

	if e != nil {
		return nil, e
	}

	return c.PagesBySpace(s.Identifier, "")
}
