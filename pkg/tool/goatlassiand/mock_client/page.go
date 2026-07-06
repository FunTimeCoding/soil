package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (c *Client) Page(identifier string) (*page.Page, error) {
	e, okay := c.pages[identifier]

	if !okay || e.deleted {
		return nil, fmt.Errorf("page not found: %s", identifier)
	}

	if e.page == nil {
		return nil, fmt.Errorf("page not found: %s", identifier)
	}

	return toPage(e.page), nil
}
