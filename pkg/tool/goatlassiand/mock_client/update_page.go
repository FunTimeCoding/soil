package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) UpdatePage(
	identifier string,
	title string,
	markdown string,
	message string,
) (*page.Page, error) {
	e, okay := c.pages[identifier]

	if !okay || e.Deleted || e.Page == nil {
		return nil, not_found.New("page", identifier)
	}

	e.Page.Title = title
	e.Page.Body.Storage.Value = markdown
	e.Page.Version.Number++
	e.Page.Version.Message = message

	return toPage(e.Page), nil
}
