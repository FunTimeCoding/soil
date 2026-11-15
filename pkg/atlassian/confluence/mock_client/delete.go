package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) Delete(pageIdentifier string) error {
	if c.Failure != nil {
		return c.Failure
	}

	if f := c.failures.Caller(); f != nil {
		return f
	}

	e, okay := c.pages[pageIdentifier]

	if !okay || e.Deleted || e.Page == nil || e.Page.Status != "current" {
		return not_found.New(constant.ConfluencePageKind, pageIdentifier)
	}

	e.Deleted = true

	return nil
}
