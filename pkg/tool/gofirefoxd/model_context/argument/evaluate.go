package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type Evaluate struct {
	TabIdentifier integer.Integer `json:"tab_id"`
	Expression    string          `json:"expression"`
}
