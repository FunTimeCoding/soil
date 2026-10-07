package issue_update

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

type Update struct {
	Before           *issue.Issue
	After            *issue.Issue
	CustomFieldNames []string
}
