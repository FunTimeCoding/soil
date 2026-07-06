package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type SendKey struct {
	SessionIdentifier string          `json:"session_id"`
	Keys              []string        `json:"keys"`
	Interval          integer.Integer `json:"interval"`
}
