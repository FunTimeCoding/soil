package convert

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func assigneeName(i *issue.Issue) string {
	if i.Raw.Fields.Assignee != nil {
		return i.Raw.Fields.Assignee.DisplayName
	}

	return ""
}
