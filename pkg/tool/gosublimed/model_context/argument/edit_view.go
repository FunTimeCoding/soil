package argument

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/request/boolean"
	"github.com/funtimecoding/soil/pkg/generative/mark/request/integer"
)

type EditView struct {
	ViewIdentifier integer.Integer `json:"view_id"`
	OldString      string          `json:"old_string"`
	NewString      string          `json:"new_string"`
	ReplaceAll     boolean.Boolean `json:"replace_all"`
}
