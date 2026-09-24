package index

import "github.com/funtimecoding/soil/pkg/crap/function"

func (i *Index) Add(f *function.Function) {
	i.Functions = append(i.Functions, f)
	i.byKey[f.Key()] = f
}
