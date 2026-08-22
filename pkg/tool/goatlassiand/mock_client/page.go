package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) Page(identifier string) (*page.Page, error) {
	e, okay := c.pages[identifier]

	if !okay || e.deleted {
		return nil, not_found.New("page", identifier)
	}

	if e.page == nil {
		return nil, not_found.New("page", identifier)
	}

	return toPage(e.page), nil
}
