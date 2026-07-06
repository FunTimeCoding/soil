package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page/page_post"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) CreatePage(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	markdown string,
) (*page.Page, error) {
	body, e := c.basic.PostV2Path(
		constant.Page,
		page_post.New(
			spaceIdentifier,
			parentIdentifier,
			title,
			page.ToStorage(markdown),
		).Encode(),
	)

	if e != nil {
		return nil, e
	}

	var result *response.Page
	notation.MustDecode(body, &result, false)

	return page.New(result, c.host), nil
}
