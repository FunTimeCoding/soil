package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (c *Client) DraftPage(identifier string) (*page.Page, error) {
	e, okay := c.pages[identifier]

	if !okay || e.deleted {
		return nil, fmt.Errorf("page not found: %s", identifier)
	}

	if e.page != nil && e.page.Status == constant.DraftStatus {
		return toPage(e.page), nil
	}

	return nil, fmt.Errorf("page not found: %s", identifier)
}
