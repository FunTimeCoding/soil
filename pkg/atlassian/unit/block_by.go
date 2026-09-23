package unit

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func blockBy(
	i *jira.Issue,
	key string,
	status string,
) {
	i.Fields.IssueLinks = append(
		i.Fields.IssueLinks,
		&jira.IssueLink{
			Type: jira.IssueLinkType{Inward: constant.JiraBlockedBy},
			InwardIssue: &jira.Issue{
				Key:    key,
				Fields: &jira.IssueFields{Status: &jira.Status{Name: status}},
			},
		},
	)
}
