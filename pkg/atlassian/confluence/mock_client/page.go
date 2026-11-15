package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) Page(identifier string) (*page.Page, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if f := c.failures.Caller(); f != nil {
		return nil, f
	}

	e, okay := c.pages[identifier]

	if !okay || e.Deleted {
		return nil, not_found.New(constant.ConfluencePageKind, identifier)
	}

	if e.Page == nil {
		return nil, not_found.New(constant.ConfluencePageKind, identifier)
	}

	return toPage(e.Page), nil
}
