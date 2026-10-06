package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) ChildPages(
	space string,
	name string,
) ([]*page.Page, error) {
	parent, e := c.PageBySpaceAndName(space, name)

	if e != nil {
		return nil, e
	}

	var children *response.Pages

	if f := c.basic.Get(
		c.basic.Base().Copy().Path(
			"%s/%s%s",
			constant.ConfluencePage,
			parent.Identifier,
			constant.ConfluenceChildren,
		).String(),
		&children,
	); f != nil {
		return nil, f
	}

	var result []*page.Page

	for _, p := range children.Results {
		v, g := c.Page(p.Identifier)

		if g != nil {
			return nil, g
		}

		result = append(result, v)
	}

	return result, nil
}
