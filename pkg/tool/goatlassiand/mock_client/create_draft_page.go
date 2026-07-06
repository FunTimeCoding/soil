package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (c *Client) CreateDraftPage(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	markdown string,
) (*page.Page, error) {
	return c.createWithStatus(
		spaceIdentifier,
		parentIdentifier,
		title,
		markdown,
		constant.DraftStatus,
	)
}
