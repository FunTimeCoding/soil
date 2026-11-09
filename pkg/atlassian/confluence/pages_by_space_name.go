package confluence

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (c *Client) PagesBySpaceName(n string) ([]*page.Page, error) {
	s, e := c.Space(n)

	if e != nil {
		return nil, e
	}

	return c.SpacePages(s.Identifier, "")
}
