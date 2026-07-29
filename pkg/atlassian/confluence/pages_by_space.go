package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) PagesBySpace(
	identifier string,
	status string,
) ([]*page.Page, error) {
	if status == "" {
		status = constant.ConfluenceCurrentStatus
	}

	if status == constant.ConfluenceDraftStatus {
		return c.draftPagesBySpace(identifier)
	}

	l := c.basic.Base().Copy().Path(
		"%s/%s%s",
		constant.ConfluenceSpace,
		identifier,
		constant.ConfluencePage,
	).Set(
		constant.ConfluenceBodyFormat,
		constant.ConfluenceStorageFormat,
	).Set(constant.ConfluenceStatus, status).String()
	var result []*response.Page

	for {
		body, e := c.basic.GetV2(l)

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

	return page.Sort(page.NewSlice(result, c.host)), nil
}
