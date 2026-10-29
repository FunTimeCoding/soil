package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) IssueByIdentifier(
	_ string,
	identifier string,
) (*response.Issue, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	i := c.issueIndex(identifier)

	if i < 0 {
		return nil, not_found.New("issue", identifier)
	}

	return c.issues[i], nil
}
