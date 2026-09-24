package index

import "github.com/funtimecoding/soil/pkg/crap/function"

func (i *Index) ByKey(key string) *function.Function {
	return i.byKey[key]
}
