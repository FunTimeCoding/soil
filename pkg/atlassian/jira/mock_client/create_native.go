package mock_client

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) CreateNative(i *jira.Issue) (*issue.Issue, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	created := *i
	created.Key = fmt.Sprintf("%s-%d", i.Fields.Project.Key, len(c.Issues)+1)
	result := issue.New(&created, issue.FixtureOption())
	c.Issues = append(c.Issues, result)

	return result, nil
}
