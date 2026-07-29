package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) ChildPages(
	space string,
	name string,
) ([]*page.Page, error) {
	parent, e := c.PageBySpaceAndName(space, name)

	if e != nil {
		return nil, e
	}

	body, f := c.basic.GetV2(
		c.basic.Base().Copy().Path(
			"%s/%s%s",
			constant.ConfluencePage,
			parent.Identifier,
			constant.ConfluenceChildren,
		).String(),
	)

	if f != nil {
		return nil, f
	}

	var children *response.Pages
	notation.MustDecode(body, &children, false)
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
