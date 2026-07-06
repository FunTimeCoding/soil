package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (c *Client) Favorites() ([]*page.Page, error) {
	results, e := c.Search("favorite=currentUser()")

	if e != nil {
		return nil, e
	}

	var result []*page.Page

	for _, r := range results {
		if r.Raw.Type == constant.PageType {
			p, f := c.Page(r.Raw.Identifier)

			if f != nil {
				return nil, f
			}

			result = append(result, p)
		}
	}

	return result, nil
}
