package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) UpdateIssue(
	_ string,
	identifier string,
	status string,
	_ string,
) (*response.Issue, error) {
	i := c.issueIndex(identifier)

	if i < 0 {
		return nil, not_found.New("issue", identifier)
	}

	if status != "" {
		c.issues[i].Status = status
	}

	return c.issues[i], nil
}
