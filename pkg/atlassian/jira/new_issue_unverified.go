package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) NewIssueUnverified(
	projectKey string,
	issueType string,
	summary string,
	description string,
) (*jira.Issue, error) {
	user, e := c.User()

	if e != nil {
		return nil, e
	}

	result := issue.RawStub()
	result.Fields.Reporter = user
	result.Fields.Project = jira.Project{Key: projectKey}
	result.Fields.Type = jira.IssueType{Name: issueType}
	result.Fields.Summary = summary
	result.Fields.Description = description

	return result, nil
}
