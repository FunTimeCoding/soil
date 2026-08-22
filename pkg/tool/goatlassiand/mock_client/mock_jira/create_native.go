package mock_jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) CreateNative(_ *jira.Issue) (*issue.Issue, error) {
	return nil, nil
}
