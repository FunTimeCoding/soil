package argument

import (
	"github.com/funtimecoding/soil/pkg/generative/mark/request/boolean"
	"github.com/funtimecoding/soil/pkg/generative/mark/request/integer"
)

type Fill struct {
	TabIdentifier integer.Integer `json:"tab_id"`
	Selector      string          `json:"selector"`
	Value         string          `json:"value"`
	Direct        boolean.Boolean `json:"direct"`
}
