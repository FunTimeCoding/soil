package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
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
			constant.ProjectName:     projectKey,
			constant.IssueTypeName:   issueType,
			constant.SummaryName:     summary,
			constant.DescriptionName: description,
		},
	)
}
