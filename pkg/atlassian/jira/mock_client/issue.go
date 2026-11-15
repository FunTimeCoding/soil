package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) Issue(key string) (*issue.Issue, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	for _, i := range c.Issues {
		if i.Key == key {
			return i, nil
		}
	}

	return nil, not_found.New("issue", key)
}
