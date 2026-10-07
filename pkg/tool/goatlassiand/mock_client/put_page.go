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

	if !okay || e.Deleted {
		return nil, not_found.New("page", identifier)
	}

	if status == constant.ConfluenceDraftStatus {
		if e.Page == nil {
			return nil, not_found.New("page", identifier)
		}

		draft := *e.Page
		draft.Title = title
		draft.Body.Storage.Value = body
		draft.Version.Number = version
		draft.Version.Message = message
		draft.Status = constant.ConfluenceDraftStatus
		e.Draft = &draft

		return toPage(e.Draft), nil
	}

	if e.Page == nil {
		return nil, not_found.New("page", identifier)
	}

	e.Page.Title = title
	e.Page.Body.Storage.Value = body
	e.Page.Version.Number = version
	e.Page.Version.Message = message
	e.Page.Status = constant.ConfluenceCurrentStatus
	e.Draft = nil

	return toPage(e.Page), nil
}
