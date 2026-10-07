package issue_update

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func New(
	before *issue.Issue,
	after *issue.Issue,
	customFieldNames []string,
) *Update {
	return &Update{
		Before:           before,
		After:            after,
		CustomFieldNames: customFieldNames,
	}
}
