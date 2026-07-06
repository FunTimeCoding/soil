package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type Navigate struct {
	TabIdentifier integer.Integer `json:"tab_id"`
	Locator       string          `json:"url"`
}
