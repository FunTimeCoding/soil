package convert

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"

type UpdateResult struct {
	Issue   *server.JiraIssue `json:"issue"`
	Changes []FieldChange     `json:"changes,omitempty"`
}
