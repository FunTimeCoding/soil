package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type CloseView struct {
	ViewIdentifier integer.Integer `json:"view_id"`
}
