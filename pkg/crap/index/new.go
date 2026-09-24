package index

import "github.com/funtimecoding/soil/pkg/crap/function"

func New(root string) *Index {
	return &Index{Root: root, byKey: map[string]*function.Function{}}
}
