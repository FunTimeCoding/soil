package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type NavigateBack struct {
	TabIdentifier integer.Integer `json:"tab_id"`
}
