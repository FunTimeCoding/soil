package argument

import "github.com/funtimecoding/soil/pkg/generative/mark/request/integer"

type SetTabColor struct {
	SessionIdentifier string          `json:"session_id"`
	Red               integer.Integer `json:"red"`
	Green             integer.Integer `json:"green"`
	Blue              integer.Integer `json:"blue"`
}
