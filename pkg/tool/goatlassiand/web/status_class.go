package web

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func statusClass(i *issue.Issue) string {
	if i.Blocked() {
		return "status-blocked"
	}

	switch i.Raw.Fields.Status.StatusCategory.Key {
	case "indeterminate":
		return "status-progress"
	case "new":
		return "status-waiting"
	}

	return ""
}
