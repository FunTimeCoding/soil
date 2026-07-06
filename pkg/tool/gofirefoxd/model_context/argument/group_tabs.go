package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type GroupTabs struct {
	TabIdentifiers  []int           `json:"tab_ids"`
	GroupIdentifier integer.Integer `json:"group_id"`
	Title           string          `json:"title"`
	Color           string          `json:"color"`
}
