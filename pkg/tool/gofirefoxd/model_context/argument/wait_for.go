package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type WaitFor struct {
	TabIdentifier integer.Integer `json:"tab_id"`
	Selector      string          `json:"selector"`
	Timeout       integer.Integer `json:"timeout"`
}
