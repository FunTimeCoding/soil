package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) DraftPages() ([]*page.Page, error) {
	l := locator.New(c.host).Base(constant.ConfluenceOldBase).Path(
		"/content",
	).Set(
		"type",
		constant.ConfluencePageType,
	).Set(
		constant.ConfluenceStatus,
		constant.ConfluenceDraftStatus,
	).Set(
		constant.ConfluenceExpand,
		"body.storage,version",
	).String()
	var result []*response.Page

	for {
		body, e := c.basic.Get(l)

		if e != nil {
			return nil, e
		}

		var s *response.Pages
		notation.MustDecode(body, &s, false)
		result = append(result, s.Results...)

		if s.Links.Next == "" {
			break
		}

		l = c.basic.Next(s.Links.Next)
	}

	return page.NewSlice(result, c.host), nil
}
