package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type Click struct {
	TabIdentifier integer.Integer `json:"tab_id"`
	Selector      string          `json:"selector"`
}
