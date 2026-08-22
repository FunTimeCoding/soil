package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) PutPage(
	identifier string,
	title string,
	body string,
	version int,
	message string,
	status string,
) (*page.Page, error) {
	e, okay := c.pages[identifier]

	if !okay || e.deleted {
		return nil, not_found.New("page", identifier)
	}

	if status == constant.ConfluenceDraftStatus {
		if e.page == nil {
			return nil, not_found.New("page", identifier)
		}

		draft := *e.page
		draft.Title = title
		draft.Body.Storage.Value = body
		draft.Version.Number = version
		draft.Version.Message = message
		draft.Status = constant.ConfluenceDraftStatus
		e.draft = &draft

		return toPage(e.draft), nil
	}

	if e.page == nil {
		return nil, not_found.New("page", identifier)
	}

	e.page.Title = title
	e.page.Body.Storage.Value = body
	e.page.Version.Number = version
	e.page.Version.Message = message
	e.page.Status = constant.ConfluenceCurrentStatus
	e.draft = nil

	return toPage(e.page), nil
}
