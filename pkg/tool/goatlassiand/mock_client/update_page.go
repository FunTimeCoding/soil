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

	if !okay || e.deleted || e.page == nil {
		return nil, not_found.New("page", identifier)
	}

	e.page.Title = title
	e.page.Body.Storage.Value = markdown
	e.page.Version.Number++
	e.page.Version.Message = message

	return toPage(e.page), nil
}
