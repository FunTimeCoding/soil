package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) AddComment(
	pageIdentifier string,
	body string,
) error {
	e, okay := c.pages[pageIdentifier]

	if !okay || e.Deleted {
		return not_found.New(constant.ConfluencePageKind, pageIdentifier)
	}

	return nil
}
