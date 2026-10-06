package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page/page_file"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page/page_post"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func (c *Client) Import(
	space string,
	parent string,
	base string,
	name string,
) (*page.Page, error) {
	f := page_file.Decode(system.ReadFile(base, name))
	s, e := c.SpaceByName(space)

	if e != nil {
		return nil, e
	}

	p, g := c.PageBySpaceAndName(space, parent)

	if g != nil {
		return nil, g
	}

	var result *response.Page

	if h := c.basic.PostV2Path(
		constant.ConfluencePage,
		page_post.New(
			s.Identifier,
			p.Identifier,
			f.Name,
			page.ToMarkup(f.Body),
		).Encode(),
		&result,
	); h != nil {
		return nil, h
	}

	return page.New(result, c.host), nil
}
