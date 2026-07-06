package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) EpicIssues(name string) ([]*issue.Issue, error) {
	return c.SearchFull("%s = %s", constant.ParentEpic, name)
}
