package service

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

type IssueUpdate struct {
	Before           *issue.Issue
	After            *issue.Issue
	CustomFieldNames []string
}
