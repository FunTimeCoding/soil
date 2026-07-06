package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (c *Client) DraftOverlay(identifier string) (*page.Page, error) {
	e, okay := c.pages[identifier]

	if !okay || e.deleted {
		return nil, fmt.Errorf("page not found: %s", identifier)
	}

	if e.draft != nil {
		return toPage(e.draft), nil
	}

	if e.page != nil {
		return toPage(e.page), nil
	}

	return nil, fmt.Errorf("page not found: %s", identifier)
}
