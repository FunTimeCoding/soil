package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
)

type IssueWithCustomFields struct {
	*server.JiraIssue
	Links        []*IssueLink           `json:"links,omitempty"`
	Comments     []*IssueComment        `json:"comments,omitempty"`
	Checklist    []*checklist_item.Item `json:"checklist,omitempty"`
	CustomFields map[string]string      `json:"custom_fields,omitempty"`
}
