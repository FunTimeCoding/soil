package option

import "github.com/funtimecoding/soil/pkg/source/index"

func (f *Fix) IndexDirectory() string {
	if f.Index == "" {
		return index.DefaultDirectory()
	}

	return f.Index
}
