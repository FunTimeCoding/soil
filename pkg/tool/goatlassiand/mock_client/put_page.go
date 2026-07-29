package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
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
		return nil, fmt.Errorf("page not found: %s", identifier)
	}

	if status == constant.ConfluenceDraftStatus {
		if e.page == nil {
			return nil, fmt.Errorf("page not found: %s", identifier)
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
		return nil, fmt.Errorf("page not found: %s", identifier)
	}

	e.page.Title = title
	e.page.Body.Storage.Value = body
	e.page.Version.Number = version
	e.page.Version.Message = message
	e.page.Status = constant.ConfluenceCurrentStatus
	e.draft = nil

	return toPage(e.page), nil
}
