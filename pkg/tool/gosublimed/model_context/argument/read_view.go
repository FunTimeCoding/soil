package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type ReadView struct {
	ViewIdentifier integer.Integer `json:"view_id"`
}
