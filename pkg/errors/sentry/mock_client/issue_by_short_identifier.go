package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
)

func (c *Client) IssueByShortIdentifier(
	_ string,
	identifier string,
) (*issue.Issue, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, i := range c.issues {
		if i.ShortIdentifier == identifier {
			return issue.New(i), nil
		}
	}

	return nil, not_found.New("issue", identifier)
}
