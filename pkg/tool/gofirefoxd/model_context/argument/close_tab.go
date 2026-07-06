package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type CloseTab struct {
	TabIdentifier integer.Integer `json:"tab_id"`
}
