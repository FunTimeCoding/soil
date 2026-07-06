package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) ChildPagesByIdentifier(
	identifier string,
) ([]*page.Page, error) {
	body, e := c.basic.GetV2(
		c.basic.Base().Copy().Path(
			"%s/%s%s",
			constant.Page,
			identifier,
			constant.Children,
		).String(),
	)

	if e != nil {
		return nil, e
	}

	var children *response.Pages
	notation.MustDecode(body, &children, false)
	var result []*page.Page

	for _, p := range children.Results {
		result = append(result, page.New(p, c.host))
	}

	return result, nil
}
