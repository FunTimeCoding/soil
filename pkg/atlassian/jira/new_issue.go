package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) NewIssue(
	projectKey string,
	issueType string,
	summary string,
	description string,
) (*jira.Issue, error) {
	p, e := c.MetaProject(projectKey)

	if e != nil {
		return nil, e
	}

	return jira.InitIssueWithMetaAndFields(
		p,
		p.GetIssueTypeWithName(issueType),
		map[string]string{
			constant.JiraProjectName:     projectKey,
			constant.JiraIssueTypeName:   issueType,
			constant.JiraSummaryName:     summary,
			constant.JiraDescriptionName: description,
		},
	)
}
