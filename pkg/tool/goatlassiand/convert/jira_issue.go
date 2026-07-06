package convert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"time"
)

func JiraIssue(i *issue.Issue) *server.JiraIssue {
	result := &server.JiraIssue{
		Key:     i.Key,
		Summary: i.Summary,
		Status:  i.Status,
		Type:    i.Type,
	}

	if i.Description != "" {
		result.Description = &i.Description
	}

	if i.Priority != "" {
		result.Priority = &i.Priority
	}

	if i.Raw.Fields.Assignee != nil &&
		i.Raw.Fields.Assignee.DisplayName != "" {
		result.Assignee = &i.Raw.Fields.Assignee.DisplayName
	}

	if len(i.Labels) > 0 {
		result.Labels = &i.Labels
	}

	if i.Create != nil && !i.Create.IsZero() {
		result.Created = new(i.Create.Format(time.RFC3339))
	}

	if i.Due != nil && !i.Due.IsZero() {
		result.Due = new(i.Due.Format(time.RFC3339))
	}

	if i.Link != "" {
		result.Link = &i.Link
	}

	return result
}
