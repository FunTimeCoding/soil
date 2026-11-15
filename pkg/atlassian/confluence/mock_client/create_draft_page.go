package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) CreateDraftPage(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	markdown string,
) (*page.Page, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return c.createWithStatus(
		spaceIdentifier,
		parentIdentifier,
		title,
		markdown,
		constant.ConfluenceDraftStatus,
	)
}
