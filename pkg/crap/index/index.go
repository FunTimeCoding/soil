package index

import "github.com/funtimecoding/soil/pkg/crap/function"

type Index struct {
	Root      string
	Functions []*function.Function
	byKey     map[string]*function.Function
}
