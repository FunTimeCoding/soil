package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) ChildPagesByIdentifier(
	identifier string,
) ([]*page.Page, error) {
	var children *response.Pages

	if e := c.basic.Get(
		c.basic.Base().Copy().Path(
			"%s/%s%s",
			constant.ConfluencePage,
			identifier,
			constant.ConfluenceChildren,
		).String(),
		&children,
	); e != nil {
		return nil, e
	}

	var result []*page.Page

	for _, p := range children.Results {
		result = append(result, page.New(p, c.host))
	}

	return result, nil
}
