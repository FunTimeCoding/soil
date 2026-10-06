package face

import "github.com/funtimecoding/soil/pkg/lint/fact"

type Set struct {
	byMethod map[string][]*fact.Interface
	known    map[string]bool
}
