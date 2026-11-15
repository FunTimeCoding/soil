package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
)

func (c *Client) IssueOption() (*option.Issue, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return issue.FixtureOption(), nil
}
