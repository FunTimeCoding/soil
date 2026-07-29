package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func priorityField(i *jira.Issue) string {
	if i.Fields != nil && i.Fields.Priority != nil {
		return i.Fields.Priority.Name
	}

	return constant.JiraDefaultPriority
}
