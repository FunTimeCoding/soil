package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
)

func (c *Client) IssueOption() (*option.Issue, error) {
	return issue.FixtureOption(), nil
}
