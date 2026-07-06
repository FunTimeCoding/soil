package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type UpdateGroup struct {
	GroupIdentifier integer.Integer `json:"group_id"`
	Title           string          `json:"title"`
	Color           string          `json:"color"`
}
