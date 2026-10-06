package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page/page_post"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) CreatePage(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	markdown string,
) (*page.Page, error) {
	var result *response.Page

	if e := c.basic.PostV2Path(
		constant.ConfluencePage,
		page_post.New(
			spaceIdentifier,
			parentIdentifier,
			title,
			page.ToStorage(markdown),
		).Encode(),
		&result,
	); e != nil {
		return nil, e
	}

	return page.New(result, c.host), nil
}
