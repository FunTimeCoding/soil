package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) AddComment(
	pageIdentifier string,
	body string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	if f := c.failures.Caller(); f != nil {
		return f
	}

	e, okay := c.pages[pageIdentifier]

	if !okay || e.Deleted {
		return not_found.New(constant.ConfluencePageKind, pageIdentifier)
	}

	return nil
}
